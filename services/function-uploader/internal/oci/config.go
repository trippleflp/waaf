package oci

import (
	"fmt"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"time"
)

func Config(fileName string, fileDigest digest.Digest) v1.Image {
	return v1.Image{
		Created:      func() *time.Time { t := time.Now().UTC(); return &t }(),
		Architecture: "amd64",
		OS:           "linux",
		Config: v1.ImageConfig{
			Env: []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			Cmd: []string{
				//"/bin/sh",
				//"-c",
				fmt.Sprintf("/%s", fileName),
			},
		},
		RootFS: v1.RootFS{
			Type: "layers",
			DiffIDs: []digest.Digest{
				fileDigest,
			},
		},
		History: []v1.History{
			{
				Created:    func() *time.Time { t := time.Now().UTC(); return &t }(),
				CreatedBy:  fmt.Sprintf("/bin/sh -c #(nop) ADD %s", fileName),
				EmptyLayer: true,
			},
			{
				Created:   func() *time.Time { t := time.Now().UTC(); return &t }(),
				CreatedBy: fmt.Sprintf("/bin/sh -c #(nop) CMD [/%s]", fileName),
			},
		},
	}
}
