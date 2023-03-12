package oci

import (
	"fmt"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func Config(fileName string, fileDigest digest.Digest) v1.Image {
	return v1.Image{
		Architecture: "amd64",
		OS:           "linux",
		Config: v1.ImageConfig{
			Env: []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			Cmd: []string{
				"/bin/sh", "-c", fmt.Sprintf("[/%s]", fileName),
			},
		},
		RootFS: v1.RootFS{
			Type: "layers",
			DiffIDs: []digest.Digest{
				fileDigest,
			},
		},
	}
}
