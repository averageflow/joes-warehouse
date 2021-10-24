package views

import (
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
			c.LinkStylesheet("/styles/bulma.min.css"),
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
