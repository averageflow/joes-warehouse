package views

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

// ErrorUploadingView will return the view to be shown when an error uploading and processing a data file occurred.
func ErrorUploadingView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Error uploading | Joe's Warehouse",
		Description: "An error occurred while uploading the file to the server, please try again.",
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
					Class("title is-2 is-danger"),
					g.Text("Error uploading"),
				),
				P(g.Text("An error occurred while uploading the file to the server, please try again.")),
				P(g.Text("Please submit a valid .json file.")),
			),
		},
	})
}

// ErrorSellingView will return the view to be shown when an error selling an article occurred.
func ErrorSellingView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Error selling | Joe's Warehouse",
		Description: "An error occurred while selling the product, please try again.",
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
					Class("title is-2 is-danger"),
					g.Text("Error selling"),
				),
				P(g.Text("An error occurred while selling the product, please try again.")),
			),
		},
	})
}

// ErrorLoadingView will return the view to be shown when an unspecified error occurred.
func ErrorLoadingView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Error loading | Joe's Warehouse",
		Description: "An error occurred while loading, please try again.",
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
					Class("title is-2 is-danger"),
					g.Text("Error loading"),
				),
				P(g.Text("An error occurred while loading, please try again.")),
			),
		},
	})
}

// SuccessUploadingView will return the view to be shown when a data file is uploaded and processed successfully.
func SuccessUploadingView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Success uploading | Joe's Warehouse",
		Description: "Uploaded file to the server successfully.",
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
					Class("title is-2 is-success"),
					g.Text("Success uploading"),
				),
				P(g.Text("Uploaded file to the server successfully.")),
			),
		},
	})
}

// SuccessSellingView will return the view to be shown when the sale of a product is successful.
func SuccessSellingView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Success selling | Joe's Warehouse",
		Description: "Sold products successfully.",
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
					Class("title is-2 is-success"),
					g.Text("Success selling"),
				),
				P(g.Text("Sold products successfully.")),
			),
		},
	})
}
