package articles

import (
	"reflect"
	"testing"
)

func TestConvertRawArticle(t *testing.T) {
	t.Parallel()

	type args struct {
		articleData []RawArticle
	}

	tests := []struct {
		name string
		args args
		want []Article
	}{
		{
			name: "test empty slice does not crash",
			args: args{articleData: nil},
			want: []Article{},
		},
		{
			name: "test conversion is successful with empty value",
			args: args{articleData: []RawArticle{
				{},
			}},
			want: []Article{
				{
					ID:    0,
					Name:  "",
					Stock: 0,
				},
			},
		},
		{
			name: "test conversion is successful",
			args: args{articleData: []RawArticle{
				{
					ID:    "1",
					Name:  "Name",
					Stock: "9",
				},
			}},
			want: []Article{
				{
					ID:    1,
					Name:  "Name",
					Stock: 9,
				},
			},
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			if got := ConvertRawArticle(tests[i].args.articleData); !reflect.DeepEqual(got, tests[i].want) {
				t.Errorf("ConvertRawArticle() = %v, want %v", got, tests[i].want)
			}
		})
	}
}

func TestConvertRawArticleFromProductFile(t *testing.T) {
	t.Parallel()

	type args struct {
		articleData []RawArticleFromProductFile
	}

	tests := []struct {
		name string
		args args
		want []ArticleProductRelation
	}{
		{
			name: "test empty slice does not crash",
			args: args{articleData: nil},
			want: []ArticleProductRelation{},
		},
		{
			name: "test conversion is successful with empty value",
			args: args{articleData: []RawArticleFromProductFile{
				{},
			}},
			want: []ArticleProductRelation{
				{
					ID:       0,
					AmountOf: 0,
				},
			},
		},
		{
			name: "test conversion is successful",
			args: args{articleData: []RawArticleFromProductFile{
				{
					ID:    "1",
					Name:  "Name",
					Stock: "9",
				},
			}},
			want: []ArticleProductRelation{
				{
					ID:       1,
					AmountOf: 9,
				},
			},
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			if got := ConvertRawArticleFromProductFile(tests[i].args.articleData); !reflect.DeepEqual(got, tests[i].want) {
				t.Errorf("ConvertRawArticleFromProductFile() = %v, want %v", got, tests[i].want)
			}
		})
	}
}
