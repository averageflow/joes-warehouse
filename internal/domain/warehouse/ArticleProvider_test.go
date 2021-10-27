package warehouse

import (
	"reflect"
	"testing"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
)

func Test_getArticlesForProduct(t *testing.T) {
	t.Parallel()

	type args struct {
		db         infrastructure.ApplicationDatabase
		productIDs []int64
	}

	tests := []struct {
		name    string
		args    args
		want    articles.ArticlesOfProductMap
		wantErr bool
	}{
		{
			name: "test getting records has no errors",
			args: args{
				db:         &infrastructure.MockApplicationDatabase{},
				productIDs: []int64{1, 2, 3},
			},
			want:    articles.ArticlesOfProductMap{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getArticlesForProducts(tt.args.db, tt.args.productIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("getArticlesForProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArticlesForProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArticles(t *testing.T) {
	t.Parallel()

	type args struct {
		db     infrastructure.ApplicationDatabase
		limit  int64
		offset int64
	}

	tests := []struct {
		name    string
		args    args
		want    *articles.ArticleResponseData
		wantErr bool
	}{
		{
			name: "test getting records does not error",
			args: args{
				db:     &infrastructure.MockApplicationDatabase{},
				limit:  100,
				offset: 0,
			},
			want: &articles.ArticleResponseData{
				Data: make(map[int64]articles.WebArticle),
				Sort: []int64{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetArticles(tt.args.db, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddArticles(t *testing.T) {
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
			name: "test adding records",
			args: args{
				db: &infrastructure.MockApplicationDatabase{},
				articleData: []articles.Article{
					{
						ID:    1,
						Name:  "Name",
						Stock: 1,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddArticles(tt.args.db, tt.args.articleData); (err != nil) != tt.wantErr {
				t.Errorf("AddArticles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddArticleProductRelation(t *testing.T) {
	t.Parallel()

	type args struct {
		db          infrastructure.ApplicationDatabase
		productID   int
		articleData []articles.ArticleProductRelation
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test adding records",
			args: args{
				db:        &infrastructure.MockApplicationDatabase{},
				productID: 1,
				articleData: []articles.ArticleProductRelation{
					{
						ID:       1,
						AmountOf: 12,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddArticleProductRelation(tt.args.db, tt.args.productID, tt.args.articleData); (err != nil) != tt.wantErr {
				t.Errorf("AddArticleProductRelation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteArticles(t *testing.T) {
	type args struct {
		db         infrastructure.ApplicationDatabase
		articleIDs []int64
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test delete records does not error",
			args: args{
				db:         &infrastructure.MockApplicationDatabase{},
				articleIDs: []int64{1, 2, 3},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteArticles(tt.args.db, tt.args.articleIDs); (err != nil) != tt.wantErr {
				t.Errorf("DeleteArticles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
