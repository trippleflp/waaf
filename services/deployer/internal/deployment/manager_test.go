package deployment

import (
	"fmt"
	"testing"
)

//func TestMarshall(t *testing.T) {
//	f := []*WaafFunction{
//		{
//			Name:  "Name",
//			Image: "Image",
//			Port:  "Port",
//		},
//	}
//	container := getContainer(f)
//	fmt.Println(container[0].Env[2].Value)
//
//	res, err := json.Marshal(f)
//
//	assert.NoError(t, err)
//	fmt.Println(string(res))
//}

func TestSmth(t *testing.T) {
	f := &WaafFunction{
		Name:  "name",
		Image: "image",
	}

	getFunctionDeployment("fnName", f)
	fmt.Printf("%v", f)
}
