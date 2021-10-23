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
