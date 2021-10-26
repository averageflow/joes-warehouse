package warehouse

import (
	"testing"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
)

func TestAddArticleStocks(t *testing.T) {
	t.Parallel()

	type args struct {
		db          infrastructure.ApplicationDatabase
		articleData []articles.Article
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test creation does not error",
			args: args{
				db: &infrastructure.MockApplicationDatabase{},
				articleData: []articles.Article{
					{
						ID:    1,
						Name:  "Name",
						Stock: 2,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddArticleStocks(tt.args.db, tt.args.articleData); (err != nil) != tt.wantErr {
				t.Errorf("AddArticleStocks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateArticlesStocks(t *testing.T) {
	t.Parallel()

	type args struct {
		db          infrastructure.ApplicationDatabase
		newStockMap map[int64]int64
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test creation does not error",
			args: args{
				db:          &infrastructure.MockApplicationDatabase{},
				newStockMap: map[int64]int64{1: 1, 2: 3},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateArticlesStocks(tt.args.db, tt.args.newStockMap); (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticlesStocks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
