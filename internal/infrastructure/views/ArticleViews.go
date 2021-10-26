package views

import (
	"fmt"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

// ArticleSubmissionView will return the view to be shown to upload data files containing articles.
func ArticleSubmissionView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Add Articles | Joe's Warehouse",
		Description: "Submit a list of new articles to be added to the warehouse.",
		Language:    "en",
		Head: []g.Node{
			faviconLinks(),
			c.LinkStylesheet(bulmaStyleSheet),
		},
		Body: []g.Node{
			applicationNavbar(),
			Main(
				Class("container section"),
				H1(
					Class("title is-2"),
					g.Text("Add Articles"),
				),
				Div(
					Class("field"),
					Label(
						Class("label"),
						For("submit-file-input"),
						g.Text("Please submit a .json file containing articles using the input below:"),
					),
				),
				submitFileForm(),
			),
		},
	})
}

// ArticleView will return the view to be shown to list articles in the warehouse.
func ArticleView(articleData *articles.ArticleResponseData) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Joe's Warehouse",
		Description: "Warehouse management software made by Joe.",
		Language:    "en",
		Head: []g.Node{
			faviconLinks(),
			c.LinkStylesheet(bulmaStyleSheet),
		},
		Body: []g.Node{
			applicationNavbar(),
			Main(
				Class("container section"),
				Div(
					H2(
						Class("title is-2 is-success"),
						g.Text("Articles in Warehouse"),
					),
					Table(
						Class("table is-striped has-text-centered"),
						THead(Tr(
							Th(g.Text("ID")),
							Th(g.Text("Name")),
							Th(g.Text("Stock")),
							Th(g.Text("Last updated")),
						)),
						TBody(articleTableBody(articleData)),
					),
				),
			),
		},
	})
}

// articleTableBody will create the article table body to be shown in the view.
func articleTableBody(articleData *articles.ArticleResponseData) g.Node {
	if articleData == nil {
		return Div()
	}

	return g.Group(g.Map(len(articleData.Sort), func(i int) g.Node {
		articleItem := articleData.Data[articleData.Sort[i]]
		return Tr(
			Td(g.Textf("%d", articleItem.ID)),
			Td(g.Text(articleItem.Name)),
			Td(g.Text(fmt.Sprintf("%d", articleItem.Stock))),
			Td(g.Text(infrastructure.EpochToHumanReadable(articleItem.UpdatedAt))),
		)
	}))
}
