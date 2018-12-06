package vinyl_test

import (
	"testing"

	"github.com/infatuation/vinyl"
)

func TestExchange(t *testing.T) {
	type typ struct {
		A, c, B string
		d       int
	}
	for i, test := range []struct {
		row  []string
		A, B string
	}{
		{[]string{"a", "1"}, "a", "1"},
		{[]string{"a", "1", "5", "7"}, "a", "1"},
		{[]string{"a", "1", "5", "7"}, "a", "1"},
		{[]string{}, "", ""},
		{[]string{"", "1"}, "", "1"},
		{[]string{"a"}, "a", ""},
	} {
		ex, err := vinyl.Exchange(typ{})
		if err != nil {
			t.Fatal(i, err)
		}
		v := ex.From(test.row).(typ)
		if v.A != test.A || v.B != test.B {
			t.Errorf("Test %d - Exhange from did not deerialize correctly - Row: %v, Val: %v", i, test.row, v)
		}
	}
}
func TestBadExchange(t *testing.T) {
	if _, err := vinyl.Exchange(struct{ A int }{}); err == nil {
		t.Errorf("Invalid exchange error condition")
	}
}
