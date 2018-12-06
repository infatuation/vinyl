package vinyl

import "reflect"

var stringType = reflect.TypeOf("")

func numExported(t reflect.Type) (c int) {
	for i := 0; i < t.NumField(); i++ {
		if isExported(t.Field(i)) {
			c++
		}
	}
	return c
}

func isExported(f reflect.StructField) bool {
	return f.PkgPath == ""
}
