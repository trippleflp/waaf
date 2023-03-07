package deployment

import (
	"context"
	"fmt"
	"k8s.io/client-go/kubernetes"
)

type deploymentBuilder struct {
	client    *kubernetes.Clientset
	name      *string
	ctx       context.Context
	functions []*waafFunction
}

func Builder(client *kubernetes.Clientset) *deploymentBuilder {
	dep := deploymentBuilder{client: client}
	return &dep
}

func (d *deploymentBuilder) SetFunctionGroupName(name string) *deploymentBuilder {
	d.name = &name
	return d
}
func (d *deploymentBuilder) SetContext(ctx context.Context) *deploymentBuilder {
	d.ctx = ctx
	return d
}

func (d *deploymentBuilder) SetFunctions(name, image string) *deploymentBuilder {
	if d.functions == nil {
		d.functions = []*waafFunction{}
	}
	d.functions = append(d.functions, &waafFunction{
		name: name, image: image,
	})
	return d
}

func (d *deploymentBuilder) Build() (*manager, error) {
	if d.client == nil {
		return nil, fmt.Errorf("client was not provided")
	}
	if d.name == nil {
		return nil, fmt.Errorf("function group name was not provided")
	}
	if d.ctx == nil {
		d.ctx = context.Background()
	}

	return getManager(*d.client, *d.name, d.ctx, d.functions), nil
}
