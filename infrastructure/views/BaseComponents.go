package views

import (
	g "github.com/maragudk/gomponents"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

func Navbar() g.Node {
	return Nav(
		Class("navbar is-transparent"),
		Div(
			Class("navbar-brand"),
			A(
				Class("navbar-item"),
				Href("/ui/products"),
				g.Text("Joe's Warehouse"),
			),
			A(

				g.Attr("onclick", `document.getElementById("navbar-menu").classList.toggle("is-active");document.getElementById("navbar-burger").classList.toggle("is-active");`),
				ID("navbar-burger"),
				Class("navbar-burger"),
				Aria("label", "menu"),
				Aria("expanded", "false"),
				Span(Aria("hidden", "true")),
				Span(Aria("hidden", "true")),
				Span(Aria("hidden", "true")),
			),
		),
		Div(
			ID("navbar-menu"),
			Class("navbar-menu"),
			Div(
				Class("navbar-start"),
				A(
					Class("navbar-item"),
					Href("/ui/articles"),
					g.Text("View articles"),
				),
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
