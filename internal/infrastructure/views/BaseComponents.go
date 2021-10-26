package views

import (
	"net/http"

	g "github.com/maragudk/gomponents"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

// submitFileForm is the re-usable file upload form.
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
				Required(),
				Class("input"),
				ID("submit-file-input"),
				Type("file"),
			),
		),
		submitFormButton(),
	)
}

// submitFormButton is a re-usable component to be used to submit any form.
func submitFormButton() g.Node {
	return Button(
		Type("submit"),
		Class("mt-4 button is-dark"),
		g.Text("Submit"),
	)
}

// navbar is the application's navigation bar as re-usable component.
func navbar() g.Node {
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
					Href("/ui/products"),
					g.Text("Products"),
				),
				A(
					Class("navbar-item"),
					Href("/ui/articles"),
					g.Text("Articles"),
				),
				A(
					Class("navbar-item"),
					Href("/ui/transactions"),
					g.Text("Transactions"),
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
