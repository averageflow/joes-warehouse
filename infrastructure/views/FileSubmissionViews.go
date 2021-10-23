package views

import (
	"net/http"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

func submitFileForm() g.Node {
	return FormEl(
		Method(http.MethodPost),
		Action(""),
		g.Attr("enctype", "multipart/form-data"),
		Div(
			Class("control"),
			Input(
				Name("fileData"),
				Accept("application/json"),
				Class("input"),
				ID("submit-file-input"),
				Type("file"),
			),
		),
		submitFormButton(),
	)
}

func submitFormButton() g.Node {
	return Button(
		Type("submit"),
		Class("mt-4 button is-dark"),
		g.Text("Submit"),
	)
}

func ArticleSubmissionView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Add Articles | Joe's Warehouse",
		Description: "Submit a list of new articles to be added to the warehouse.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet("/styles/bulma.min.css"),
		},
		Body: []g.Node{
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

func ProductSubmissionView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Add Products | Joe's Warehouse",
		Description: "Submit a list of new products to be added to the warehouse.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet("/styles/bulma.min.css"),
		},
		Body: []g.Node{
			Main(
				Class("container has-text-justified p-6"),
				H1(
					Class("title is-2"),
					g.Text("Add Products"),
				),
				Div(
					Class("field"),
					Label(
						Class("label"),
						For("submit-file-input"),
						g.Text("Please submit a .json file containing products using the input below:"),
					),
				),
				submitFileForm(),
			),
		},
	})
}
