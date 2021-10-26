package products

import (
	"math"
	"sort"
)

// CollectProductIDs will return the product IDs from a map of products.
func CollectProductIDs(products map[int64]WebProduct) []int64 {
	var result []int64

	for i := range products {
		result = append(result, products[i].ID)
	}

	return result
}

// CollectProductIDs will return the product IDs from a map of product and stocks in a sale request.
func CollectProductIDsForSell(products map[int64]int64) []int64 {
	var result []int64

	for i := range products {
		result = append(result, i)
	}

	return result
}

// ProductAmountInStock will calculate the amount of stock for a product
// based on the stock of the different articles that constitute a product.
func ProductAmountInStock(product WebProduct) int64 {
	if len(product.Articles) == 0 {
		// products should always consist of articles
		// this edge case must be still handled, and thus we return max infinity
		return 0
	}

	var amounts []float64

	for i := range product.Articles {
		article := product.Articles[i]
		if article.AmountOf > article.Stock {
			// if we need more parts than are in stock then we immediately stop
			// the calculation and return a 0
			return 0
		}

		ratio := float64(article.Stock / article.AmountOf)
		amounts = append(amounts, ratio)
	}

	// by knowing the smallest amount of times we can use existing articles in stock
	// we can deduce the maximum amount of products we can sell
	sort.Float64s(amounts)

	return int64(math.Floor(amounts[0]))
}
