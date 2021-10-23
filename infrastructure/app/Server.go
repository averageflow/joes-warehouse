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

	"github.com/averageflow/joes-warehouse/infrastructure"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	gracefulShutdownRequestGraceSeconds = 10
)

type ApplicationState struct {
	Handler    infrastructure.ApplicationHTTPHandler
	HTTPServer *http.Server
	DB         *pgxpool.Pool
	Config     *ApplicationConfig
}

type ApplicationServer struct {
	State ApplicationState
}

type ApplicationServerResponse struct {
	Message       string `json:"message,omitempty"`
	Error         string `json:"error,omitempty"`
	UnixTimestamp int64  `json:"unix_timestamp"`
}

func NewApplicationServer(userOptions *ApplicationState) *ApplicationServer {
	ConfigSetup("config", ".")

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

		//db.Query()
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

func (s *ApplicationServer) registerHandlers() {
	s.State.Handler.Use(gin.Logger(), gin.Recovery())

	s.State.Handler.StaticFile("/styles/bulma.min.css", "../../web/styles/bulma.min.css")

	s.State.Handler.Handle(http.MethodGet, "/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ui/products")
	})

	uiGroup := s.State.Handler.Group("/ui")

	// HTML views
	uiGroup.Handle(http.MethodGet, "", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ui/products")
	})
	uiGroup.Handle(http.MethodGet, "/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ui/products")
	})

	uiGroup.Handle(http.MethodGet, "/products", s.productViewHandler())
	uiGroup.Handle(http.MethodGet, "/articles", s.articleViewHandler())
	uiGroup.Handle(http.MethodGet, "/products/file-submission", s.addProductsFromFileViewHandler())
	uiGroup.Handle(http.MethodGet, "/articles/file-submission", s.addArticlesFromFileViewHandler())
	// Form submissions
	uiGroup.Handle(http.MethodPost, "/articles/file-submission", s.addDataFromFileHandler(infrastructure.ItemTypeArticle))
	uiGroup.Handle(http.MethodPost, "/products/file-submission", s.addDataFromFileHandler(infrastructure.ItemTypeProduct))

	headlessGroup := s.State.Handler.Group("/api")
	headlessGroup.Handle(http.MethodGet, "/products", s.getProductsHandler())
	headlessGroup.Handle(http.MethodPost, "/products", s.addProductsHandler())
	headlessGroup.Handle(http.MethodGet, "/articles", s.getArticlesHandler())
	headlessGroup.Handle(http.MethodPost, "/articles", s.addArticlesHandler())

	// s.State.Handler.Handle(http.MethodPatch, "/products/:id", s.modifyProductHandler())
	// s.State.Handler.Handle(http.MethodDelete, "/products/:id", s.deleteProductHandler())
	// s.State.Handler.Handle(http.MethodPatch, "/articles/:id", s.modifyArticleHandler())
	// s.State.Handler.Handle(http.MethodDelete, "/articles/:id", s.deleteArticleHandler())
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
