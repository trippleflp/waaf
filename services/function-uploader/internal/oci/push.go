package oci

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

func (b *builder) pushBlob(fileBytes []byte, digest digest.Digest) error {
	location, err := b.obtainSessionId()
	if err != nil {
		return err
	}

	uri, err := url.Parse(*location)
	if err != nil {
		return err
	}
	//locationPath := locationParsed.Path
	//uri, err := url.Parse(locationPath)
	//if err != nil {
	//	return err
	//}
	fmt.Println(string(digest))
	q := uri.Query()
	q.Add("digest", digest.String())
	uri.RawQuery = q.Encode()
	req := &http.Request{
		URL: uri,
		Header: map[string][]string{
			"Content-Length": {strconv.Itoa(len(fileBytes))},
			"Content-Type":   {"application/octet-stream"},
		},
		Body:   io.NopCloser(bytes.NewReader(fileBytes)),
		Method: "PUT",
	}

	fmt.Printf("Calling: %s\n", req.URL)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 201 {
		return errors.New("Status not 201 but " + res.Status)
	}

	return nil
}

func (b *builder) pushManifest(fileBytes []byte) (string, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/%s", b.registryURL, path.Join("v2", b.name, "manifests", b.tag)))
	if err != nil {
		return "", err
	}

	req := &http.Request{
		URL: uri,
		Header: map[string][]string{
			"Content-Type": {v1.MediaTypeImageManifest},
		},
		Method: "PUT",
		//Host:   b.registryURL,
		Body: io.NopCloser(bytes.NewReader(fileBytes)),
	}

	fmt.Printf("Calling: %s\n", req.URL)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 201 {
		return "", errors.New("Status not 201 but " + res.Status)
	}
	return res.Header.Get("Location"), nil
}

func (b *builder) obtainSessionId() (*string, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/%s/", b.registryURL, path.Join("v2", b.name, "blobs", "uploads")))
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		URL:    uri,
		Method: "POST",
		//Host:   b.registryURL,
	}
	fmt.Printf("Calling: %s\n", req.URL)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	location := res.Header.Get("Location")
	if location == "" {
		return nil, errors.New("sessionid is empty")
	}
	fmt.Printf("Location: %s\n", location)
	return &location, nil
}
