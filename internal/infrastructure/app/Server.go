package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/averageflow/joes-warehouse/internal/domain/warehouse"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	// gracefulShutdownRequestGraceSeconds is the time the application waits to close
	// any currently processing HTTP requests will gracefully shutting down.
	gracefulShutdownRequestGraceSeconds = 10
)

const (
	defaultPaginationLimit         = 100
	defaultFrontendPaginationLimit = 500
)

// ApplicationState is the application's state in a centralized manner.
type ApplicationState struct {
	Handler    infrastructure.ApplicationHTTPHandler
	HTTPServer *http.Server
	DB         infrastructure.ApplicationDatabase
	Config     *ApplicationConfig
}

// ApplicationServer is the core application's handler,
// the contact surface with outside world.
type ApplicationServer struct {
	State ApplicationState
}

// ApplicationServerResponse is the standard response the server returns when using JSON.
type ApplicationServerResponse struct {
	Message       string `json:"message,omitempty"`
	Error         string `json:"error,omitempty"`
	UnixTimestamp int64  `json:"unix_timestamp"`
}

// NewApplicationServer will return a fully configured ApplicationServer.
func NewApplicationServer(userOptions *ApplicationState) *ApplicationServer {
	http.DefaultClient.Timeout = 30 * time.Second

	state := userOptions
	if state == nil {
		state = &ApplicationState{}
	}

	if state.Config == nil {
		state.Config = GetConfig()
	}

	if strings.EqualFold(state.Config.ApplicationMode, gin.ReleaseMode) {
		// the application mode should be set before initializing the HTTP handler
		// so that it takes effect and routes produce less verbose output
		gin.SetMode(gin.ReleaseMode)
	}

	state.Handler = gin.New()

	if state.HTTPServer == nil {
		state.HTTPServer = &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 60 * time.Second,
			IdleTimeout:  60 * time.Second,
			Addr:         ":7000",
			Handler:      state.Handler,
		}
	}

	if state.DB == nil {
		db, err := pgxpool.Connect(context.Background(), state.Config.DatabaseConnection)
		if err != nil {
			log.Fatalln(fmt.Sprintf("Unable to connect to database: %v\n", err))
			os.Exit(1)
		}

		state.DB = db
	}

	srv := ApplicationServer{
		State: ApplicationState{
			HTTPServer: state.HTTPServer,
			Handler:    state.Handler,
			Config:     state.Config,
			DB:         state.DB,
		},
	}

	srv.registerHandlers()

	return &srv
}

// registerHandlers will prepare the HTTP handler and associate routes to handling methods.
func (s *ApplicationServer) registerHandlers() {
	s.State.Handler.Use(gin.Logger(), gin.Recovery())

	// API calls
	headlessGroup := s.State.Handler.Group("/api")
	headlessGroup.Use(s.authTokenMiddleware())

	headlessGroup.Handle(http.MethodGet, "/products", s.getProductsHandler())
	headlessGroup.Handle(http.MethodPost, "/products", s.addProductsHandler())
	headlessGroup.Handle(http.MethodDelete, "/products/:id", s.deleteProductHandler())
	headlessGroup.Handle(http.MethodPatch, "/products/sell", s.sellProductsHandler())

	headlessGroup.Handle(http.MethodGet, "/articles", s.getArticlesHandler())
	headlessGroup.Handle(http.MethodPost, "/articles", s.addArticlesHandler())
	headlessGroup.Handle(http.MethodDelete, "/articles/:id", s.deleteArticleHandler())

	headlessGroup.Handle(http.MethodGet, "/transactions", s.getTransactionsHandler())

	// UI related routes
	s.State.Handler.Static("/assets/favicon", fmt.Sprintf("%s/assets/favicon", s.State.Config.WebAssetLocation))

	uiGroup := s.State.Handler.Group("/ui")

	// UI related redirects
	s.State.Handler.Handle(http.MethodGet, "/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ui/products")
	})
	uiGroup.Handle(http.MethodGet, "", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ui/products")
	})
	uiGroup.Handle(http.MethodGet, "/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ui/products")
	})

	// HTML views
	uiGroup.Handle(http.MethodGet, "/products", s.productViewHandler())
	uiGroup.Handle(http.MethodGet, "/articles", s.articleViewHandler())
	uiGroup.Handle(http.MethodGet, "/transactions", s.transactionViewHandler())

	uiGroup.Handle(http.MethodGet, "/products/file-submission", s.addProductsFromFileViewHandler())
	uiGroup.Handle(http.MethodGet, "/articles/file-submission", s.addArticlesFromFileViewHandler())

	// Form submissions
	uiGroup.Handle(http.MethodPost, "/articles/file-submission", s.addArticlesFromFileHandler())
	uiGroup.Handle(http.MethodPost, "/products/file-submission", s.addProductsFromFileHandler())
	uiGroup.Handle(http.MethodPost, "/products/sell", s.sellProductFormHandler())
}

// getPaginationDetails will extract the pagination details from the URL parameters in a GET endpoint.
func (s *ApplicationServer) getPaginationDetails(c *gin.Context) warehouse.PaginationInQueryParams {
	var paginationDetails warehouse.PaginationInQueryParams

	// ignore the error since the values for pagination will default to 0
	// making the application more resilient
	_ = c.BindQuery(&paginationDetails)

	if paginationDetails.Limit < 1 {
		// default pagination limit if it was not specified or invalid
		paginationDetails.Limit = defaultPaginationLimit
	}

	if paginationDetails.Offset < 1 {
		// default pagination offset to 0 if it was not specified or invalid
		paginationDetails.Offset = 0
	}

	return paginationDetails
}

// TerminationSignalWatcher will wait for interrupt signal to gracefully shutdown
// the server with a timeout of x seconds.
func TerminationSignalWatcher(srv *http.Server) {
	// Make a channel to receive operating system signals.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
	log.Println("Shutting down server, because of received signal..")

	// The context is used to inform the server it has x seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(
		context.Background(),
		gracefulShutdownRequestGraceSeconds*time.Second,
	)

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	defer cancel()

	log.Println("Server exiting..")
}
