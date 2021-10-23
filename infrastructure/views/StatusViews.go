package views

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

func ErrorUploadingView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Error uploading | Joe's Warehouse",
		Description: "An error occurred while uploading the file to the server, please try again.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet("/styles/bulma.min.css"),
		},
		Body: []g.Node{
			Navbar(),
			Main(
				Class("container has-text-justified p-6"),
				H1(
					Class("title is-2 is-danger"),
					g.Text("Error uploading"),
				),
				P(g.Text("An error occurred while uploading the file to the server, please try again.")),
				P(g.Text("Please submit a valid .json file.")),
			),
		},
	})
}

func SuccessUploadingView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Success uploading | Joe's Warehouse",
		Description: "Uploaded file to the server successfully.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet("/styles/bulma.min.css"),
		},
		Body: []g.Node{
			Navbar(),
			Main(
				Class("container has-text-justified p-6"),
				H1(
					Class("title is-2 is-success"),
					g.Text("Success uploading"),
				),
				P(g.Text("Uploaded file to the server successfully.")),
			),
		},
	})
}

func HomeView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Joe's Warehouse",
		Description: "Warehouse management software made by Joe.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet("/styles/bulma.min.css"),
		},
		Body: []g.Node{
			Navbar(),
			Main(
				Class("container has-text-justified p-6"),
				Div(
					H2(
						Class("title is-2 is-success"),
						g.Text("Products"),
					),
				),
				Div(
					H2(
						Class("title is-2 is-success"),
						g.Text("Articles"),
					),
				),
			),
		},
	})
}

func Navbar() g.Node {
	return Nav(
		Class("navbar is-transparent"),
		Div(
			Class("navbar-brand"),
			A(
				Class("navbar-item"),
				Href("/"),
				g.Text("Joe's Warehouse"),
			),
		),
		Div(
			Class("navbar-menu"),
			Div(
				Class("navbar-start"),
				A(
					Class("navbar-item"),
					Href("/ui/products/file-submission"),
					g.Text("Add products"),
				),
				A(
					Class("navbar-item"),
					Href("/ui/articles/file-submission"),
					g.Text("Add articles"),
				),
			),
		),
	)
}
