package utils

import "fmt"

// Values is a generic type for argument in a query
type Values map[string]string

// BuildUpdateArgs for exec query
func BuildUpdateArgs(id string, args *Values) (string, []interface{}) {
	ret := []interface{}{id}
	expr := ""

	i := 2
	for k, v := range *args {
		if expr != "" {
			expr = expr + ", "
		}
		expr = expr + fmt.Sprintf("%s=$%d", k, i)
		ret = append(ret, v)
		i = i + 1
	}
	return expr, ret
}
