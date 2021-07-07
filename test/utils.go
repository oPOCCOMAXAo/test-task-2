package test

import (
	"reflect"
	"testing"
)

// if expected != got test t will fail with descriptive message
func Check(t *testing.T, label string, expected interface{}, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error in %s: expected %#v, got %#v", label, expected, got)
	}
}
