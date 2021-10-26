package products

import (
	"reflect"
	"testing"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
)

func TestCollectProductIDs(t *testing.T) {
	t.Parallel()

	type args struct {
		products map[int64]WebProduct
	}

	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test empty slice does not crash",
			args: args{products: nil},
			want: nil,
		},
		{
			name: "test conversion is successful",
			args: args{products: map[int64]WebProduct{
				1: {ID: 1},
			}},
			want: []int64{1},
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			if got := CollectProductIDs(tests[i].args.products); !reflect.DeepEqual(got, tests[i].want) {
				t.Errorf("CollectProductIDs() = %v, want %v", got, tests[i].want)
			}
		})
	}
}

func TestCollectProductIDsForSell(t *testing.T) {
	t.Parallel()

	type args struct {
		products map[int64]int64
	}

	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test empty slice does not crash",
			args: args{products: nil},
			want: nil,
		},
		{
			name: "test conversion is successful",
			args: args{products: map[int64]int64{
				1: 123,
			}},
			want: []int64{1},
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			if got := CollectProductIDsForSell(tests[i].args.products); !reflect.DeepEqual(got, tests[i].want) {
				t.Errorf("CollectProductIDsForSell() = %v, want %v", got, tests[i].want)
			}
		})
	}
}

func TestProductAmountInStock(t *testing.T) {
	t.Parallel()

	type args struct {
		product WebProduct
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test empty slice returns 0",
			args: args{product: WebProduct{}},
			want: 0,
		},
		{
			name: "test enough items return positive stock",
			args: args{product: WebProduct{
				Articles: map[int64]articles.ArticleOfProduct{
					1: {Stock: 8, AmountOf: 2},
				},
			}},
			want: 4,
		},
		{
			name: "test enough items return smallest positive stock",
			args: args{product: WebProduct{
				Articles: map[int64]articles.ArticleOfProduct{
					1: {Stock: 8, AmountOf: 2},
					2: {Stock: 2, AmountOf: 2},
				},
			}},
			want: 1,
		},
		{
			name: "test enough items return 0 if one article missing",
			args: args{product: WebProduct{
				Articles: map[int64]articles.ArticleOfProduct{
					1: {Stock: 8, AmountOf: 2},
					2: {Stock: 0, AmountOf: 2},
				},
			}},
			want: 0,
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			if got := ProductAmountInStock(tests[i].args.product); got != tests[i].want {
				t.Errorf("ProductAmountInStock() = %v, want %v", got, tests[i].want)
			}
		})
	}
}
