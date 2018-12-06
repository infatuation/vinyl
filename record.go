// Package vinyl creates Records from structs inferring schema from exported fields via reflect. This package can make working with the stdlib "encoding/csv" more ergonomic by supplying a translation layer between structs and csv encoding that doesn't require extra annotation or codegen. All string marhsaling of labels and values is done via fmt.Sprint
package vinyl

import (
	"fmt"
	"reflect"
)

// Record is a generic type for holding interface types
type Record struct {
	v reflect.Value
	t reflect.Type
	//	len int
}

// New will create a new record from the supplied interface
func New(v interface{}) Record {
	t := reflect.TypeOf(v)
	return Record{
		v: reflect.ValueOf(v),
		t: t,
		//	len: numExported(t),
	}
}

// Values will translate the record into a string slice.
// The values of the slice correspond to the records labels index.
func (rec Record) Values() []string {
	row := make([]string, numExported(rec.t))
	var k int
	for i := 0; i < rec.t.NumField(); i++ {
		if isExported(rec.t.Field(i)) {
			row[k] = fmt.Sprint(rec.v.Field(i))
			k++
		}
	}
	return row
}

// Dict returns the records as values with the labels as keys
// Dict is the slowest as it also reflects labels from the records fields
// The Formatters will be applied to the keys in the order they are supplied
func (rec Record) Dict(apply ...Formatter) map[string]string {
	row := make(map[string]string, numExported(rec.t))
	for i := 0; i < rec.v.NumField(); i++ {
		if isExported(rec.t.Field(i)) {
			v := rec.v.Field(i)
			k := rec.t.Field(i).Name
			for _, a := range apply {
				k = a(k)
			}
			row[k] = fmt.Sprint(v)
		}
	}
	return row
}

func (rec Record) fields() []string {
	h := make([]string, numExported(rec.t))
	var k int
	for i := 0; i < rec.t.NumField(); i++ {
		if isExported(rec.t.Field(i)) {
			h[k] = rec.t.Field(i).Name
			k++
		}
	}
	return h
}
