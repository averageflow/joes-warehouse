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
			Class("field has-addons"),
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
			Div(
				Class("control"),
				submitFormButton(),
			),
		),
	)
}

// submitFormButton is a re-usable component to be used to submit any form.
func submitFormButton() g.Node {
	return Button(
		Type("submit"),
		Class("button is-dark"),
		g.Text("Submit"),
	)
}

// navbar is the application's navigation bar as re-usable component.
func applicationNavbar() g.Node {
	return Nav(
		Class("navbar is-transparent has-background-white-ter"),
		Div(
			Class("navbar-brand"),
			A(
				Class("navbar-item"),
				Href("/ui/products"),
				Img(
					Src("/assets/favicon/apple-icon-120x120.png"),
				),
			),
			A(
				g.Attr("onclick", `document.getElementById("navbar-menu").classList.toggle("is-active");document.getElementById("navbar-burger").classList.toggle("is-active");`),
				ID("navbar-burger"),
				Class("navbar-burger has-text-weight-bold"),
				Aria("label", "menu"),
				Aria("expanded", "false"),
				Span(Aria("hidden", "true")),
				Span(Aria("hidden", "true")),
				Span(Aria("hidden", "true")),
			),
		),
		Div(
			ID("navbar-menu"),
			Class("navbar-menu has-background-white-ter"),
			Div(
				Class("navbar-start"),
				A(
					Class("navbar-item has-text-weight-bold"),
					Href("/ui/products"),
					g.Text("Products"),
				),
				A(
					Class("navbar-item has-text-weight-bold"),
					Href("/ui/articles"),
					g.Text("Articles"),
				),
				A(
					Class("navbar-item has-text-weight-bold"),
					Href("/ui/transactions"),
					g.Text("Transactions"),
				),
				A(
					Class("navbar-item has-text-weight-bold"),
					Href("/ui/products/file-submission"),
					g.Text("+ Products"),
				),
				A(
					Class("navbar-item has-text-weight-bold"),
					Href("/ui/articles/file-submission"),
					g.Text("+ Articles"),
				),
			),
		),
	)
}

func faviconLinks() g.Node {
	return g.Raw(`
		<link rel="apple-touch-icon" sizes="57x57" href="/assets/apple-icon-57x57.png">
		<link rel="apple-touch-icon" sizes="60x60" href="/assets/favicon/apple-icon-60x60.png">
		<link rel="apple-touch-icon" sizes="72x72" href="/assets/favicon/apple-icon-72x72.png">
		<link rel="apple-touch-icon" sizes="76x76" href="/assets/favicon/apple-icon-76x76.png">
		<link rel="apple-touch-icon" sizes="114x114" href="/assets/favicon/apple-icon-114x114.png">
		<link rel="apple-touch-icon" sizes="120x120" href="/assets/favicon/apple-icon-120x120.png">
		<link rel="apple-touch-icon" sizes="144x144" href="/assets/favicon/apple-icon-144x144.png">
		<link rel="apple-touch-icon" sizes="152x152" href="/assets/favicon/apple-icon-152x152.png">
		<link rel="apple-touch-icon" sizes="180x180" href="/assets/favicon/apple-icon-180x180.png">
		<link rel="icon" type="image/png" sizes="192x192"  href="/assets/favicon/android-icon-192x192.png">
		<link rel="icon" type="image/png" sizes="32x32" href="/assets/favicon/favicon-32x32.png">
		<link rel="icon" type="image/png" sizes="96x96" href="/assets/favicon/favicon-96x96.png">
		<link rel="icon" type="image/png" sizes="16x16" href="/assets/favicon/favicon-16x16.png">
		<link rel="manifest" href="/assets/favicon/manifest.json">
		<meta name="msapplication-TileColor" content="#ffffff">
		<meta name="msapplication-TileImage" content="/assets/favicon/ms-icon-144x144.png">
		<meta name="theme-color" content="#333">
	`)
}
