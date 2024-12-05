package postgres

import (
	"fmt"
	"strings"
)

func placeholders(n, start int) string {
	if n < 1 {
		return ""
	}
	placeholders := make([]string, n)
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+start)
	}
	return strings.Join(placeholders, ", ")
}
