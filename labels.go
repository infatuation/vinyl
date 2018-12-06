package vinyl

import (
	"strings"

	"github.com/fatih/camelcase"
)

// Formatter is a func to rename record labels
type Formatter func(string) string

// Labels represent the schema of the record and are inferred from the exported Fields of the supplied type
// Labels are deterministic in their names and order as defined by the supplied type
// Formatters will be applied in the order they are supplied
func Labels(v interface{}, apply ...Formatter) []string {
	fields := New(v).fields()
	for i, h := range fields {
		for _, a := range apply {
			fields[i] = a(h)
		}
	}
	return fields
}

// SnakeFormat is a formatter to convert labels into snake_case
func SnakeFormat(v string) string {
	return snaked(v)
}
func snaked(h string) string {
	return strings.ToLower(strings.Join(camelcase.Split(h), "_"))
}
