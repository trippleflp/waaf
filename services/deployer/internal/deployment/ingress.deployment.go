package deployment

import (
	"fmt"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getIngressDeployment(functionGroupName string, namespace string) *v1.Ingress {
	groupStack := fmt.Sprintf("%s-stack", functionGroupName)
	ingressDeployment := &v1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      fmt.Sprintf("%s-ingress", groupStack),
			Annotations: map[string]string{
				"kubernetes.io/ingress.class":                 "nginx",
				"nginx.ingress.kubernetes.io/rewrite-target":  "/$2",
				"nginx.ingress.kubernetes.io/proxy-body-size": "8m",
			},
		},
		Spec: v1.IngressSpec{
			Rules: []v1.IngressRule{
				{
					IngressRuleValue: v1.IngressRuleValue{
						HTTP: &v1.HTTPIngressRuleValue{
							Paths: []v1.HTTPIngressPath{
								{
									Path:     fmt.Sprintf("/waaf/fn/%s(/|$)(.*)", functionGroupName),
									PathType: (func() *v1.PathType { path := v1.PathTypePrefix; return &path })(),
									Backend: v1.IngressBackend{
										Service: &v1.IngressServiceBackend{
											Name: fmt.Sprintf("%s-service", groupStack),
											Port: v1.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return ingressDeployment
}
