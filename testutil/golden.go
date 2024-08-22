package testutil

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

var updateGolden = flag.Bool("update", false, "update golden files")

func Golden(t *testing.T, data []byte, ext string) {
	t.Helper()

	if *updateGolden {
		UpdateGolden(t, data, ext)
		t.SkipNow()
	}
	CompareGolden(t, data, ext)
}

func UpdateGolden(t *testing.T, data []byte, ext string) {
	t.Helper()

	t.Log("updating golden archive")
	file := filepath.Join("testdata", "golden", t.Name()+ext)
	if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
		t.Fatalf("creating golden archive directory: %s", err)
	}
	if err := os.WriteFile(file, data, 0o644); err != nil {
		t.Fatalf("writing golden archive: %s", err)
	}
}

func CompareGolden(t *testing.T, data []byte, ext string) {
	t.Helper()

	file := filepath.Join("testdata", "golden", t.Name()+ext)
	golden, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("reading golden archive: %s", err)
	}
	AssertNoDiff(t, golden, data)
}
