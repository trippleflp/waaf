package deployment

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshall(t *testing.T) {
	f := []*waafFunction{
		{
			name:  "name",
			image: "image",
			port:  "port",
		},
	}
	container := getContainer(f)
	fmt.Println(container[0].Env[2].Value)

	res, err := json.Marshal(*(f[0]))
	assert.NoError(t, err)
	fmt.Println(string(res))
}
