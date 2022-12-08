package errors

import (
	"strings"
)

// funBaseName removes the path prefix component of a function's name reported by func.Name().
func funBaseName(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
