package vinyl_test

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"testing"
	"time"

	"github.com/infatuation/vinyl"
)

func BenchmarkCSVEncoding(b *testing.B) {
	type test struct {
		A string
		B int
		d time.Time
		C time.Time
	}
	buf := bytes.NewBuffer(nil)
	w := csv.NewWriter(buf)
	for i := 0; i < b.N; i++ {
		r := vinyl.New(test{fmt.Sprint(i), i, time.Now(), time.Now()})
		if err := w.Write(r.Values()); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLabels(b *testing.B) {
	test := struct {
		A string
		B int
		C time.Time
	}{"a", 2, time.Now()}
	for i := 0; i < b.N; i++ {
		lbl := vinyl.Labels(test)
		_ = lbl
	}
}

func BenchmarkValues(b *testing.B) {
	test := struct {
		A string
		B int
		C time.Time
	}{"a", 2, time.Now()}
	rec := vinyl.New(test)
	for i := 0; i < b.N; i++ {
		s := rec.Values()
		_ = s
	}
}

func BenchmarkDict(b *testing.B) {
	test := struct {
		A string
		B int
		C time.Time
	}{"a", 2, time.Now()}
	rec := vinyl.New(test)
	for i := 0; i < b.N; i++ {
		s := rec.Dict()
		_ = s
	}
}

func BenchmarkNewRecord(b *testing.B) {
	test := struct {
		A string
		B int
		C time.Time
	}{"a", 2, time.Now()}
	for i := 0; i < b.N; i++ {
		rec := vinyl.New(test)
		_ = rec
	}
}

func TestNewRecord(t *testing.T) {
	now := time.Now()
	for i, test := range []struct {
		v     interface{}
		slice []string
	}{
		{
			struct {
				A string
				B int
				C time.Time
			}{"a", 1, now},
			[]string{"a", "1", fmt.Sprint(now)},
		},
	} {
		rec := vinyl.New(test.v)
		s := rec.Values()
		if len(s) != len(test.slice) {
			t.Errorf("Test %d - Invalid Record Slice - Want: %v, Got: %v", i, test.slice, s)
		}
		for i, v := range test.slice {
			if s[i] != v {
				t.Errorf("Test %d - Invalid Record Slice - Want: %v, Got: %v", i, test.slice, s)
			}
		}
	}
}
