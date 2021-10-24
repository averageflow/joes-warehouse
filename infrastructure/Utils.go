package infrastructure

import (
	"fmt"
	"strings"
)

func IntSliceToCommaSeparatedString(data []int64) string {
	tmp := make([]string, len(data))

	for i := range data {
		tmp[i] = fmt.Sprintf("%d", data[i])
	}

	return strings.Join(tmp, ", ")
}
