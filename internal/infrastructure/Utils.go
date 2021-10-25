package infrastructure

import (
	"fmt"
	"strings"
)

// IntSliceToCommaSeparatedString will convert a slice of int64 items into
// a comma separated string.
func IntSliceToCommaSeparatedString(data []int64) string {
	tmp := make([]string, len(data))

	for i := range data {
		tmp[i] = fmt.Sprintf("%d", data[i])
	}

	return strings.Join(tmp, ", ")
}
