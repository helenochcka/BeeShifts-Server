package repositories

import (
	"fmt"
	"strings"
)

func buildFilter(field []interface{}, fieldName string, conditions *[]string, args *[]interface{}) {
	if len(field) > 0 {
		*conditions = append(*conditions, fmt.Sprintf("%s IN (%s)", fieldName, placeholders(len(field), len(*args)+1)))
		for _, arg := range field {
			*args = append(*args, arg)
		}
	}
}

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
