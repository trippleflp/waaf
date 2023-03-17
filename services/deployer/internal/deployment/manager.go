package deployment

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	v12 "k8s.io/api/networking/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	networkv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type WaafFunction struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	PodName     string `json:"podName"`
	ServiceName string `json:"service"`
}

type manager struct {
	client            kubernetes.Clientset
	functionGroupName string
	deploymentClient  v1.DeploymentInterface
	servicesClient    corev1.ServiceInterface
	ingressClient     networkv1.IngressInterface
	ctx               context.Context
	functions         []*WaafFunction
	namespace         string
}

func getManager(client kubernetes.Clientset, name string, ctx context.Context, functions []*WaafFunction) (*manager, error) {
	namespace := createNamespace(client, name, ctx)

	configMap, err := client.CoreV1().ConfigMaps("waaf").Get(ctx, "nginx-stack-config", metav1.GetOptions{})
	if err != nil {
		log.Err(err).Msg("Could not load configmap")
		return nil, err
	}
	copyConfigMap := &apiv1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		Data: configMap.Data,
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-stack-config", name),
			Namespace: namespace,
		},
	}
	_, err = client.CoreV1().ConfigMaps(namespace).Create(ctx, copyConfigMap, metav1.CreateOptions{})
	if err != nil {
		if strings.Compare(err.Error(), fmt.Sprintf("configmaps \"%s-stack-config\" already exists", name)) != 0 {
			log.Print("Could not create copy of configmap")
			return nil, err
		}
	}

	m := manager{
		client:            client,
		functionGroupName: name,
		deploymentClient:  client.AppsV1().Deployments(namespace),
		servicesClient:    client.CoreV1().Services(namespace),
		ingressClient:     client.NetworkingV1().Ingresses(namespace),
		ctx:               ctx,
		functions:         functions,
		namespace:         namespace,
	}
	return &m, nil
}

func createNamespace(clientset kubernetes.Clientset, name string, ctx context.Context) string {
	isDeleted := false
	ns := fmt.Sprintf("waaf-ns-%s", name)
	nsSpec := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"nsKey": ns,
			},
			Name: ns,
		},
	}
	err := clientset.CoreV1().Namespaces().Delete(ctx, ns, metav1.DeleteOptions{})
	if err != nil {
		if strings.Compare(err.Error(), fmt.Sprintf("namespaces \"%s\" not found", ns)) == 0 {
			isDeleted = true
		} else {
			log.Err(err).Msg("Deletion of namespace failed")

		}
	}
	watcher, err := clientset.CoreV1().Namespaces().Watch(ctx, metav1.ListOptions{
		LabelSelector:  fmt.Sprintf("nsKey=%s", ns),
		TimeoutSeconds: func() *int64 { i := int64(200); return &i }(),
		Watch:          true,
	})
	if err != nil {
		log.Err(err).Msg("Watching of namespace deletion failed")
		isDeleted = true
	}
	if !isDeleted {
		for event := range watcher.ResultChan() {
			log.Print(event.Type)
			if event.Type == watch.Deleted {
				watcher.Stop()
				break
			}
			log.Err(err).Msg(fmt.Sprintf("Namespace %s is still termintating", ns))
		}

		log.Printf(fmt.Sprintf("Namespace %s is deleted", ns))
	}

	_, err = clientset.CoreV1().Namespaces().Create(ctx, nsSpec, metav1.CreateOptions{})
	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("was not able to create namespace %s", ns))
	}
	return ns
}

func (m *manager) DeployPod() error {
	for _, deployment := range getNginxDeployment(m.functionGroupName, m.functions, m.namespace) {
		log.Printf("deploying pod: %s", deployment.Name)

		result, err := m.deploymentClient.Create(m.ctx, deployment, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		fmt.Printf("%s created.\n", result.Name)
	}

	return nil
}

func (m *manager) DeployService() error {
	_, err := m.servicesClient.Create(m.ctx, getServiceDeployment(m.functionGroupName, m.namespace), metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Service for %s function group created\n", m.functionGroupName)
	for _, function := range m.functions {

		log.Printf("deplyoing service for function: %s", function.Name)
		_, err := m.servicesClient.Create(m.ctx, getFunctionServiceDeployment(function, m.namespace), metav1.CreateOptions{})
		if err != nil {
			return err
		}
		fmt.Printf("Service for function %s created\n", function.Name)

	}

	return err
}

func (m *manager) DeployIngress() (*v12.Ingress, error) {
	result, err := m.ingressClient.Create(m.ctx, getIngressDeployment(m.functionGroupName, m.namespace), metav1.CreateOptions{})
	fmt.Printf("Ingress for %s function group started\n", m.functionGroupName)
	return result, err
}

func (m *manager) DeployAll() error {
	log.Printf("Creating functions: %v", m.functions)

	err := m.DeployPod()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		//return err
	}
	err = m.DeployService()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		//return err
	}
	_, err = m.DeployIngress()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	return err
}
