package converter

import (
	"fmt"
	"strconv"
)

func boolish(ish interface{}) (result bool, err error) {
	if ish == nil {
		return
	}

	switch notIsh := ish.(type) {
	case string:
		value, err := strconv.ParseBool(notIsh)
		if err != nil {
			return false, fmt.Errorf("error parsing bool: %w, value: %t", err, value)
		}
	case int:
		if ish == 1 {
			result = true
			return
		}
	case bool:
		result = notIsh
		return
	}
	return
}
