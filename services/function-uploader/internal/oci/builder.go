package oci

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/go-digest"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type builder struct {
	filePath    string
	filename    string
	err         error
	registryURL string
	tag         string
	name        string
}

func Builder() *builder {
	return &builder{}
}

func (b *builder) SetFile(filePath string) *builder {
	b.filePath = filePath
	b.filename = filepath.Base(filePath)
	return b
}

func (b *builder) SetRegistryUrl(url string) *builder {
	b.registryURL = url
	return b
}

func (b *builder) SetName(name string) *builder {
	b.name = strings.ToLower(name)
	return b
}

func (b *builder) SetTag(tag string) *builder {
	b.tag = tag
	return b
}

func (b *builder) Build() (string, error) {
	if b.err != nil {
		return "", b.err
	}

	tarPath, err := CreateTarball(filepath.Dir(b.filePath), b.filePath)
	if err != nil {
		return "", err
	}

	tarFile, err := os.Open(*tarPath)
	if err != nil {
		return "", err
	}
	defer tarFile.Close()

	tarBytes, err := io.ReadAll(tarFile)
	tarDigest := digest.FromBytes(tarBytes)

	layer := fileDetails{
		size:   int64(len(tarBytes)),
		digest: tarDigest,
	}

	config := Config(b.filename, tarDigest)
	configBytes, err := json.Marshal(config)
	if err != nil {
		return "", err
	}

	configLayer := fileDetails{
		size:   int64(len(configBytes)),
		digest: digest.FromBytes(configBytes),
	}
	manifest := buildManifest(configLayer, layer)
	manifestBytes, err := json.Marshal(manifest)
	if err != nil {
		return "", err
	}

	log.Print("Pushing config")
	err = b.pushBlob(configBytes, configLayer.digest)
	if err != nil {
		return "", err
	}
	log.Print("Pushing layer")
	err = b.pushBlob(tarBytes, layer.digest)
	if err != nil {
		return "", err
	}

	log.Print("Pushing manifest")
	_, err = b.pushManifest(manifestBytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", b.name, b.tag), nil
}
