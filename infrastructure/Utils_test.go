package infrastructure

import "testing"

func TestIntSliceToCommaSeparatedString(t *testing.T) {
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
