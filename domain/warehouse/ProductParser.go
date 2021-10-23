package warehouse

import (
	"fmt"
	"strings"

	"github.com/averageflow/joes-warehouse/infrastructure"
)

func ConvertRawProduct(products []infrastructure.RawProduct) []infrastructure.Product {
	result := make([]infrastructure.Product, len(products))

	for i := range products {
		result[i] = infrastructure.Product{
			Name: products[i].Name,
		}
	}

	return result
}

func IntSliceToCommaSeparatedString(data []int64) string {
	tmp := make([]string, len(data))

	for i := range data {
		tmp[i] = fmt.Sprintf("%d", data[i])
	}

	return strings.Join(tmp, ", ")
}

func CollectProductIDs(products map[string]infrastructure.WebProduct) []int64 {
	var result []int64

	for i := range products {
		result = append(result, products[i].ID)
	}
	return result
}

func CollectProductIDsToUniqueIDs(products map[string]infrastructure.WebProduct) map[int64]string {
	result := make(map[int64]string)

	for i := range products {
		result[products[i].ID] = products[i].UniqueID
	}

	return result
}
