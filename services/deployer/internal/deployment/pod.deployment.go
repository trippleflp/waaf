package deployment

import (
	"encoding/json"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"strconv"
	"strings"
)

func getNginxDeployment(functionGroupName string, functions []*WaafFunction) *appsv1.Deployment {
	groupStack := fmt.Sprintf("%s-stack", functionGroupName)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-deployment", groupStack),
			Labels: map[string]string{
				"app": groupStack,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": groupStack,
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": groupStack,
					},
				},
				Spec: apiv1.PodSpec{
					Containers:    getContainer(functions),
					RestartPolicy: apiv1.RestartPolicyAlways,
					Volumes: []apiv1.Volume{
						{Name: "nginx-conf", VolumeSource: apiv1.VolumeSource{
							ConfigMap: &apiv1.ConfigMapVolumeSource{
								LocalObjectReference: apiv1.LocalObjectReference{Name: "nginx.conf"},
								Items: []apiv1.KeyToPath{
									{Key: "nginx.conf", Path: "nginx.conf"},
								},
							},
						}},
						{
							Name: "log", VolumeSource: apiv1.VolumeSource{
								EmptyDir: &apiv1.EmptyDirVolumeSource{},
							},
						},
						{
							Name: "nginx-stack-config",
							VolumeSource: apiv1.VolumeSource{
								ConfigMap: &apiv1.ConfigMapVolumeSource{
									LocalObjectReference: apiv1.LocalObjectReference{Name: "nginx-stack-config"},
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment
}

func getContainer(functions []*WaafFunction) []apiv1.Container {
	currentPort := 8001
	container := []apiv1.Container{
		{
			Name:  "nginx",
			Image: "nginx",
			Ports: []apiv1.ContainerPort{
				{
					ContainerPort: 80,
				},
			},
			Env: []apiv1.EnvVar{
				{Name: "GROUP_ID", Value: "some_random_uuid"},
				{Name: "NGINX_ENVSUBST_OUTPUT_DIR", Value: "/etc/nginx"},
			},
			VolumeMounts: []apiv1.VolumeMount{
				{Name: "log", MountPath: "/var/log/nginx"},
				// TODO js path
				{Name: "nginx-stack-config", MountPath: "/etc/nginx/templates/nginx.conf.template", SubPath: "nginx.conf"},
				{Name: "nginx-stack-config", MountPath: "/etc/nginx/index.js", SubPath: "index.js"},
			},
		},
	}
	s := "'["
	for _, function := range functions {
		function.Port = strconv.Itoa(currentPort)
		container = append(container, apiv1.Container{
			Name:  function.Name,
			Image: function.Image,
			Env: []apiv1.EnvVar{
				{
					Name:  "PORT",
					Value: function.Port,
				},
			},
		})
		bytes, _ := json.Marshal(*function)
		s = fmt.Sprintf("%s%s,", s, string(bytes))
		currentPort += currentPort
	}

	s = strings.TrimSuffix(s, ",")
	s = fmt.Sprintf("%s]'", s)
	//bytes, _ := json.Marshal(functions)
	container[0].Env = append(container[0].Env, apiv1.EnvVar{Name: "FUNCTION_DATA", Value: s})

	return container
}
