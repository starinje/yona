package helpers

import (
	"fmt"
	"strconv"
)

func ParseID(idStr string) (uint, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format")
	}
	return uint(id), nil
}
