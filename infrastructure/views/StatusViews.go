package views

import (
	"fmt"
	"net/http"

	"github.com/averageflow/joes-warehouse/domain/articles"
	"github.com/averageflow/joes-warehouse/domain/products"
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
			navbar(),
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

func ErrorSellingView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Error selling | Joe's Warehouse",
		Description: "An error occurred while selling the product, please try again.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet("/styles/bulma.min.css"),
		},
		Body: []g.Node{
			navbar(),
			Main(
				Class("container has-text-justified p-6"),
				H1(
					Class("title is-2 is-danger"),
					g.Text("Error selling"),
				),
				P(g.Text("An error occurred while selling the product, please try again.")),
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
			navbar(),
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

func SuccessSellingView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Success selling | Joe's Warehouse",
		Description: "Sold products successfully.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet("/styles/bulma.min.css"),
		},
		Body: []g.Node{
			navbar(),
			Main(
				Class("container has-text-justified p-6"),
				H1(
					Class("title is-2 is-success"),
					g.Text("Success selling"),
				),
				P(g.Text("Sold products successfully.")),
			),
		},
	})
}

func ProductView(products map[int64]products.WebProduct, sortProducts []int64) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Joe's Warehouse",
		Description: "Warehouse management software made by Joe.",
		Language:    "en",
		Head: []g.Node{
			c.LinkStylesheet("/styles/bulma.min.css"),
		},
		Body: []g.Node{
			navbar(),
			Main(
				Class("container has-text-justified p-6"),
				Div(
					H2(
						Class("title is-2 is-success"),
						g.Text("Products"),
					),
					Table(
						Class("table is-striped"),
						THead(Tr(
							Th(g.Text("Name")),
							Th(g.Text("Price")),
							Th(g.Text("Stock")),
							Th(),
						)),
						TBody(
							g.Group(g.Map(len(sortProducts), func(i int) g.Node {
								return Tr(
									Td(g.Text(products[sortProducts[i]].Name)),
									Td(g.Text(fmt.Sprintf("%.2f", products[sortProducts[i]].Price))),
									Td(g.Text(fmt.Sprintf("%d", products[sortProducts[i]].AmountInStock))),
									Td(
										FormEl(
											Method(http.MethodPost),
											Action("/ui/products/sell"),
											g.Attr("enctype", "application/x-www-form-urlencoded"),
											Input(
												Type("hidden"),
												Class("is-hidden"),
												Required(),
												Name("productID"),
												Value(fmt.Sprintf("%d", sortProducts[i])),
												ReadOnly(),
											),
											Div(
												Class("control is-flex-desktop is-flex-tablet"),
												Input(
													Class("input is-small"),
													Required(),
													Name("amount"),
													Type("number"),
													Min("0"),
													Max(fmt.Sprintf("%d", products[sortProducts[i]].AmountInStock)),
												),
												Button(
													Type("submit"),
													Class("button is-dark is-small"),
													g.Text("Sell"),
												),
											),
										),
									),
								)
							})),
						),
					),
				),
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
			c.LinkStylesheet("/styles/bulma.min.css"),
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
