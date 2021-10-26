package infrastructure

import (
	"reflect"
	"testing"
)

func TestIntSliceToCommaSeparatedString(t *testing.T) {
	t.Parallel()

	type args struct {
		data []int64
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test empty slice",
			args: args{data: nil},
			want: "",
		},
		{
			name: "test empty slice",
			args: args{data: []int64{1}},
			want: "1",
		},
		{
			name: "test empty slice",
			args: args{data: []int64{1, 2, 3}},
			want: "1, 2, 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntSliceToCommaSeparatedString(tt.args.data); got != tt.want {
				t.Errorf("IntSliceToCommaSeparatedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendIfMissingInt64(t *testing.T) {
	t.Parallel()

	type args struct {
		slice []int64
		i     int64
	}

	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test does not append",
			args: args{
				slice: []int64{1},
				i:     1,
			},
			want: []int64{1},
		},
		{
			name: "test does append",
			args: args{
				slice: []int64{1, 2},
				i:     123,
			},
			want: []int64{1, 2, 123},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendIfMissingInt64(tt.args.slice, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendIfMissingInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
