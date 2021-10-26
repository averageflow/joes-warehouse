package views

import (
	"github.com/averageflow/joes-warehouse/internal/domain/transactions"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"

	// Dot import is used here to avoid having to make the code unreadable
	// with so many references to html
	// nolint
	. "github.com/maragudk/gomponents/html"
)

// ArticleView will return the view to be shown to list articles in the warehouse.
func TransactionView(transactionData *transactions.TransactionResponse) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       "Joe's Warehouse",
		Description: "Warehouse management software made by Joe.",
		Language:    "en",
		Head: []g.Node{
			faviconLinks(),
			c.LinkStylesheet(bulmaStyleSheet),
		},
		Body: []g.Node{
			applicationNavbar(),
			Main(
				Class("container section"),
				Div(
					H2(
						Class("title is-2 is-success"),
						g.Text("Transactions"),
					),
					Table(
						Class("table is-striped has-text-centered"),
						THead(Tr(
							Th(g.Text("ID")),
							Th(g.Text("ProductID")),
							Th(g.Text("Product")),
							Th(g.Text("Amount")),
							Th(g.Text("Created")),
						)),
						TBody(transactionTableBody(transactionData)),
					),
				),
			),
		},
	})
}

// articleTableBody will create the article table body to be shown in the view.
func transactionTableBody(transactionData *transactions.TransactionResponse) g.Node {
	if transactionData == nil {
		return nil
	}

	var tableRows []g.Node

	for i := range transactionData.Sort {
		transactionItems := transactionData.Data[transactionData.Sort[i]]

		for j := range transactionItems {
			transactionItem := transactionItems[j]
			tableRows = append(tableRows, Tr(
				g.If(j == 0, Td(g.Textf("%d", transactionItem.ID))),
				g.If(j != 0, Td()),
				Td(g.Textf("%d", transactionItem.ProductID)),
				Td(g.Text(transactionItem.ProductName)),
				Td(g.Textf("%d", transactionItem.ProductAmount)),
				Td(g.Textf(infrastructure.EpochToHumanReadable(transactionItem.CreatedAt))),
			))
		}
	}

	return g.Group(tableRows)
}
