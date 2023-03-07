package deployment

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type waafFunction struct {
	name  string `json:"name"`
	image string
	port  string `json:"port"`
}

type manager struct {
	client            kubernetes.Clientset
	functionGroupName string
	deploymentClient  v1.DeploymentInterface
	ctx               context.Context
	functions         []*waafFunction
}

func getManager(client kubernetes.Clientset, name string, ctx context.Context, functions []*waafFunction) *manager {
	m := manager{
		client:            client,
		functionGroupName: name,
		deploymentClient:  client.AppsV1().Deployments(apiv1.NamespaceDefault),
		ctx:               ctx,
		functions:         functions,
	}
	return &m
}

func (m *manager) DeployNginx() (*appsv1.Deployment, error) {
	result, err := m.deploymentClient.Create(m.ctx, getNginxDeployment(m.functionGroupName, m.functions), metav1.CreateOptions{})
	fmt.Printf("Function %s with image %s started with port %s\n", m.functions[0].name, m.functions[0].image, m.functions[0].port)

	return result, err
}
