package views

import (
	"fmt"
	"net/http"

	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

// ProductSubmissionView will return the view to be shown to upload data files containing products.
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

// ProductView will return the view to be shown with a list of products in the warehouse.
func ProductView(productData *products.ProductResponseData) g.Node {
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
							Th(g.Text("ID")),
							Th(g.Text("Name")),
							Th(g.Text("Price")),
							Th(g.Text("Stock")),
							Th(),
							Th(g.Text("Last updated")),
						)),
						TBody(productTableBody(productData)),
					),
				),
			),
		},
	})
}

// productTableBody will create the product table body to be shown in the view.
func productTableBody(productData *products.ProductResponseData) g.Node {
	if productData == nil {
		return Div()
	}

	return g.Group(g.Map(len(productData.Sort), func(i int) g.Node {
		productItem := productData.Data[productData.Sort[i]]
		return Tr(
			Td(g.Textf("%d", productItem.ID)),
			Td(g.Text(productItem.Name)),
			Td(g.Textf("€ %.2f", productItem.Price)),
			Td(g.Textf("%d", productItem.AmountInStock)),
			Td(sellProductForm(productItem.ID, productItem.AmountInStock)),
			Td(g.Text(infrastructure.EpochToHumanReadable(productItem.UpdatedAt))),
		)
	}))
}

// sellProductForm is the re-usable form used to submit a sell product request.
func sellProductForm(productID, amountInStock int64) g.Node {
	return FormEl(
		Method(http.MethodPost),
		Action("/ui/products/sell"),
		g.Attr("enctype", "application/x-www-form-urlencoded"),
		Input(
			Type("hidden"),
			Class("is-hidden"),
			Required(),
			Name("productID"),
			Value(fmt.Sprintf("%d", productID)),
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
				Max(fmt.Sprintf("%d", amountInStock)),
			),
			Button(
				Type("submit"),
				Class("button is-dark is-small"),
				g.Text("Sell"),
			),
		),
	)
}
