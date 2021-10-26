package infrastructure

import (
	"fmt"
	"strings"
	"time"
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

// EpochToHumanReadable will return a RFC822 format date time string from a Unix epoch.
func EpochToHumanReadable(epoch int64) string {
	return time.Unix(epoch, 0).Format(time.RFC822)
}

// AppendIfMissingInt64 will append to an int64 slice if the wanted element is missing.
func AppendIfMissingInt64(slice []int64, i int64) []int64 {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}

	return append(slice, i)
}
