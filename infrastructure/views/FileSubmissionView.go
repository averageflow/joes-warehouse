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
		Head:        []g.Node{},
		Body: []g.Node{
			Main(
				Class("container has-text-justified"),
				Label(
					g.Text("Please submit a file containing articles using the input below:"),
					Input(
						Type("file"),
					),
				),
			),
		},
	})
}
