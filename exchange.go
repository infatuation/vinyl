package vinyl

import (
	"errors"
	"reflect"
)

type exchange struct {
	t reflect.Type
}

// Exchange can read records from a row and translate them into the supplied concrete type. Only types with all exported string fields are supported.
//   type t stuct{A, B, C string}
//   recs, _ := vinyl.Exchange(t{})
//   row, _ := csvreader.Read()
//   v := recs.From(row).(t)
func Exchange(v interface{}) (ex exchange, err error) {
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if isExported(f) && f.Type != stringType { // TODO support UnmarshalText
			return ex, errors.New("Exchange only supports types with all string fields")
		}
	}
	return exchange{t}, nil
}

// From swaps a row for an interface backed by the concrete applied when opening the exchange
// If the exported field count of the concrete type is longer than the row, the row is truncated
func (ex exchange) From(row []string) interface{} {
	v := reflect.New(ex.t).Elem()
	var k int
	for i := 0; i < v.NumField() && k < len(row); i++ {
		if isExported(ex.t.Field(i)) {
			v.Field(i).SetString(row[k])
			k++
		}
	}
	return v.Interface()
}
