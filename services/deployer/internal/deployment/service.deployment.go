package deployment

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getServiceDeployment(functionGroupName string, namespace string) *corev1.Service {
	groupStack := fmt.Sprintf("%s-stack", functionGroupName)
	serviceDeployment := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      fmt.Sprintf("%s-service", groupStack),
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{Port: 80},
			},
			Selector: map[string]string{"app": groupStack},
		},
	}
	return serviceDeployment
}

func getFunctionServiceDeployment(function *WaafFunction, namespace string) *corev1.Service {
	serviceDeployment := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      function.ServiceName,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{Port: 8080},
			},
			Selector: map[string]string{"app": function.PodName},
		},
	}
	return serviceDeployment
}
