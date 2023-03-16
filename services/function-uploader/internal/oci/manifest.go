package oci

import (
	"github.com/opencontainers/go-digest"
	"github.com/opencontainers/image-spec/specs-go"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type fileDetails struct {
	size   int64
	digest digest.Digest
}

func buildManifest(config, layer fileDetails) v1.Manifest {
	return v1.Manifest{
		Versioned: specs.Versioned{SchemaVersion: 2},
		MediaType: v1.MediaTypeImageManifest,
		Config: v1.Descriptor{
			MediaType: v1.MediaTypeImageConfig,
			Digest:    config.digest,
			Size:      config.size,
		},
		Layers: []v1.Descriptor{{
			//MediaType: v1.MediaTypeImageLayer,
			MediaType: v1.MediaTypeImageLayer,
			Digest:    layer.digest,
			Size:      layer.size},
		},
		Annotations: map[string]string{
			"module.wasm.image/variant": "compat-smart",
		},
	}

}
