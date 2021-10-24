package products

import (
	"math"
	"sort"
)

func ConvertRawProduct(products []RawProduct) []Product {
	result := make([]Product, len(products))

	for i := range products {
		result[i] = Product{
			Name: products[i].Name,
		}
	}

	return result
}

func CollectProductIDs(products map[int64]WebProduct) []int64 {
	var result []int64

	for i := range products {
		result = append(result, products[i].ID)
	}

	return result
}

func CollectProductIDsForSell(products map[int64]int64) []int64 {
	var result []int64

	for i := range products {
		result = append(result, i)
	}

	return result
}

func ProductAmountInStock(product WebProduct) int64 {
	var amounts []float64

	for i := range product.Articles {
		if product.Articles[i].AmountOf > product.Articles[i].Stock {
			// if we need more parts than are in stock then we immediately stop
			// the calculation and return a 0
			return 0
		}

		ratio := float64(product.Articles[i].Stock / product.Articles[i].AmountOf)
		amounts = append(amounts, ratio)
	}

	sort.Float64s(amounts)

	return int64(math.Floor(amounts[0]))
}
