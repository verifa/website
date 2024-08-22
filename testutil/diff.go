package testutil

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func AssertNoDiff[T any](t *testing.T, expected, actual T) {
	t.Helper()
	if diff := Diff(actual, expected); diff != "" {
		t.Fatal(Callers(), diff)
	}
}

// Diff compares two items and returns a human-readable diff string. If the
// items are equal, the string is empty.
func Diff[T any](expected, actual T, opts ...cmp.Option) string {
	// nolint: gocritic
	oo := append(
		opts,
		cmp.Exporter(func(reflect.Type) bool { return true }),
		cmpopts.EquateEmpty(),
	)

	diff := cmp.Diff(actual, expected, oo...)
	if diff != "" {
		return "\n-want +got\n" + diff
	}

	return ""
}
