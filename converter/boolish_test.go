package converter

import "testing"

func TestBoolish(t *testing.T) {
	data := map[interface{}]bool{
		true:    true,
		1:       true,
		"false": false,
		32:      false,
		nil:     false,
		"0":     false,
		0:       false,
		false:   false,
		"":      false,
	}
	for val, expected := range data {
		ish, _ := boolish(val)
		if expected != ish {
			t.Errorf("Parsing %q, expected %v got %v", val, expected, ish)
		}
	}
}
