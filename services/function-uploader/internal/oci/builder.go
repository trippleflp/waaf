package oci

import (
	"encoding/json"
	"github.com/opencontainers/go-digest"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
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
	file, err := os.Open(filePath)
	if err != nil {
		b.err = err
		return b
	}
	defer file.Close()
	if b.err != nil {
		return b
	}
	b.filename = file.Name()
	return b
}

func (b *builder) SetRegistryUrl(url string) *builder {
	b.registryURL = url
	return b
}

func (b *builder) SetName(name string) *builder {
	b.name = name
	return b
}

func (b *builder) SetTag(tag string) *builder {
	b.tag = tag
	return b
}

func (b *builder) Build() error {
	if b.err != nil {
		return b.err
	}

	tarPath, err := CreateTarball(filepath.Dir(b.filename), b.filename)
	if err != nil {
		return err
	}

	tarFile, err := os.Open(*tarPath)
	if err != nil {
		return err
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
		return err
	}

	configLayer := fileDetails{
		size:   int64(len(configBytes)),
		digest: digest.FromBytes(configBytes),
	}
	manifest := buildManifest(configLayer, layer)
	manifestBytes, err := json.Marshal(manifest)
	if err != nil {
		return err
	}

	log.Print("Pushing config")
	err = b.pushBlob(configBytes, configLayer.digest)
	if err != nil {
		return err
	}
	log.Print("Pushing layer")
	err = b.pushBlob(tarBytes, layer.digest)
	if err != nil {
		return err
	}

	log.Print("Pushing manifest")
	err = b.pushManifest(manifestBytes)
	if err != nil {
		return err
	}
	return nil
}
