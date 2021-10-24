package views

import (
	"fmt"
	"net/http"

	"github.com/averageflow/joes-warehouse/internal/domain/products"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

func ProductSubmissionView() g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Add Products | Joe's Warehouse",
		Description: "Submit a list of new products to be added to the warehouse.",
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

func ProductView(products map[int64]products.WebProduct, sortProducts []int64) g.Node {
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
