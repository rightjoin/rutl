package conv

import (
	"strconv"
)

func IntOr(source string, deflt int) int {
	out, err := strconv.Atoi(source)
	if err != nil {
		return deflt
	}
	return out
}
