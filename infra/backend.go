package infra

import "github.com/golingon/lingon/pkg/terra"

var _ terra.Backend = (*GCSBackend)(nil)

type GCSBackend struct {
	Bucket string `hcl:"bucket"`
	Prefix string `hcl:"prefix"`
}

func (g *GCSBackend) BackendType() string {
	return "gcs"
}
