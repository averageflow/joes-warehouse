package warehouse

import "github.com/averageflow/joes-warehouse/infrastructure"

func ConvertRawProduct(products []infrastructure.RawProductModel) []infrastructure.ProductModel {
	result := make([]infrastructure.ProductModel, len(products))

	for i := range products {
		result[i] = infrastructure.ProductModel{
			Name: products[i].Name,
		}
	}

	return result
}
