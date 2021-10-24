package views

import (
	"fmt"

	"github.com/averageflow/joes-warehouse/domain/articles"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

func ArticleSubmissionView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Add Articles | Joe's Warehouse",
		Description: "Submit a list of new articles to be added to the warehouse.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet(bulmaStyleSheet),
		},
		Body: []g.Node{
			navbar(),
			Main(
				Class("container has-text-justified p-6"),
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

func ArticleView(articles map[int64]articles.WebArticle, sortArticles []int64) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Joe's Warehouse",
		Description: "Warehouse management software made by Joe.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet(bulmaStyleSheet),
		},
		Body: []g.Node{
			navbar(),
			Main(
				Class("container has-text-justified p-6"),
				Div(
					H2(
						Class("title is-2 is-success"),
						g.Text("Articles"),
					),
					Table(
						Class("table is-striped"),
						THead(Tr(
							Th(g.Text("Name")),
							Th(g.Text("Stock")),
						)),
						TBody(
							g.Group(g.Map(len(sortArticles), func(i int) g.Node {
								return Tr(
									Td(g.Text(articles[sortArticles[i]].Name)),
									Td(g.Text(fmt.Sprintf("%d", articles[sortArticles[i]].Stock))),
								)
							})),
						),
					),
				),
			),
		},
	})
}
