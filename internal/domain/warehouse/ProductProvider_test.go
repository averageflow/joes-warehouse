package warehouse

import (
	"reflect"
	"testing"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
)

func TestGetFullProductResponse(t *testing.T) {
	t.Parallel()

	type args struct {
		db     infrastructure.ApplicationDatabase
		limit  int64
		offset int64
	}

	tests := []struct {
		name    string
		args    args
		want    *products.ProductResponseData
		wantErr bool
	}{
		{
			name: "test get records empty warehouse does not error",
			args: args{
				db:     &infrastructure.MockApplicationDatabase{},
				limit:  30,
				offset: 0,
			},
			want:    &products.ProductResponseData{Data: make(map[int64]products.WebProduct), Sort: []int64{}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFullProductResponse(tt.args.db, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullProductResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFullProductResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFullProductsByID(t *testing.T) {
	t.Parallel()

	type args struct {
		db               infrastructure.ApplicationDatabase
		wantedProductIDs []int64
	}

	tests := []struct {
		name    string
		args    args
		want    *products.ProductResponseData
		wantErr bool
	}{
		{
			name: "test get records by id does not error",
			args: args{
				db:               &infrastructure.MockApplicationDatabase{},
				wantedProductIDs: []int64{1, 2, 3, 4},
			},
			want:    &products.ProductResponseData{Data: make(map[int64]products.WebProduct), Sort: []int64{}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFullProductsByID(tt.args.db, tt.args.wantedProductIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullProductsByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFullProductsByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddProducts(t *testing.T) {
	t.Parallel()

	type args struct {
		db          infrastructure.ApplicationDatabase
		productData []products.RawProduct
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test add records does not error",
			args: args{
				db: &infrastructure.MockApplicationDatabase{},
				productData: []products.RawProduct{
					{
						Name:  "product",
						Price: 123,
						Articles: []articles.RawArticleFromProductFile{
							{
								ID:    "2",
								Name:  "Name",
								Stock: "2",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddProducts(tt.args.db, tt.args.productData); (err != nil) != tt.wantErr {
				t.Errorf("AddProducts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSellProducts(t *testing.T) {
	t.Parallel()

	type args struct {
		db             infrastructure.ApplicationDatabase
		wantedProducts map[int64]int64
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test sell products does not error",
			args: args{
				db:             &infrastructure.MockApplicationDatabase{},
				wantedProducts: map[int64]int64{1: 1, 2: 2, 3: 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SellProducts(tt.args.db, tt.args.wantedProducts); (err != nil) != tt.wantErr {
				t.Errorf("SellProducts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteProducts(t *testing.T) {
	type args struct {
		db         infrastructure.ApplicationDatabase
		productIDs []int64
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test deletion does not error",
			args: args{
				db:         &infrastructure.MockApplicationDatabase{},
				productIDs: []int64{1, 2, 3},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteProducts(tt.args.db, tt.args.productIDs); (err != nil) != tt.wantErr {
				t.Errorf("DeleteProducts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
