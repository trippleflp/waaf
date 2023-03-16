package deployment

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

func getNginxDeployment(functionGroupName string, functions []*WaafFunction, namespace string) []*appsv1.Deployment {
	var functionDeployments []*appsv1.Deployment

	for _, function := range functions {
		functionDeployments = append(functionDeployments, getFunctionDeployment(functionGroupName, function))
	}

	servicesBytes, _ := json.Marshal(functions)

	configName := fmt.Sprintf("%s-stack-config", functionGroupName)

	groupStack := fmt.Sprintf("%s-stack", functionGroupName)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      fmt.Sprintf("%s-deployment", groupStack),
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
					Containers: []apiv1.Container{
						{
							Name:  "nginx",
							Image: "nginx:stable-alpine",
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
							Env: []apiv1.EnvVar{
								{Name: "GROUP_ID", Value: "some_random_uuid"},
								{Name: "NGINX_ENVSUBST_OUTPUT_DIR", Value: "/etc/nginx"},
								{Name: "FUNCTION_DATA", Value: base64.StdEncoding.EncodeToString(servicesBytes)},
								{Name: "NAMESPACE", Value: namespace},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{Name: "log", MountPath: "/var/log/nginx"},
								{Name: configName, MountPath: "/etc/nginx/templates/nginx.conf.template", SubPath: "nginx.conf"},
								{Name: configName, MountPath: "/etc/nginx/index.js", SubPath: "index.js"},
							},
						},
					},
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
							Name: configName,
							VolumeSource: apiv1.VolumeSource{
								ConfigMap: &apiv1.ConfigMapVolumeSource{
									LocalObjectReference: apiv1.LocalObjectReference{Name: configName},
								},
							},
						},
					},
				},
			},
		},
	}

	return append(functionDeployments, deployment)
}

func getFunctionDeployment(functionGroupName string, function *WaafFunction) *appsv1.Deployment {
	podName := fmt.Sprintf("%s-stack-fn-%s", functionGroupName, function.Name)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-deployment", podName),
			Annotations: map[string]string{
				"module.wasm.image/variant": "compat-smart",
			},
			Labels: map[string]string{
				"app": podName,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": podName,
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"module.wasm.image/variant": "compat-smart",
					},
					Labels: map[string]string{
						"app": podName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							ImagePullPolicy: apiv1.PullAlways,
							Name:            function.Name,
							Image:           function.Image,
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 8080,
								},
							},
							Env: []apiv1.EnvVar{
								{Name: "PORT", Value: "8080"},
							},
						},
					},
					RuntimeClassName: func() *string { s := "crun"; return &s }(),
					RestartPolicy:    apiv1.RestartPolicyAlways,
				},
			},
		},
	}
	function.PodName = podName
	function.ServiceName = fmt.Sprintf("%s-service", podName)
	return deployment
}

//func getContainer(functions []*WaafFunction) []apiv1.Container {
//	//currentPort := 8001
//	container := []apiv1.Container{
//		{
//			Name:  "nginx",
//			Image: "nginx",
//			Ports: []apiv1.ContainerPort{
//				{
//					ContainerPort: 80,
//				},
//			},
//			Env: []apiv1.EnvVar{
//				{Name: "GROUP_ID", Value: "some_random_uuid"},
//				{Name: "NGINX_ENVSUBST_OUTPUT_DIR", Value: "/etc/nginx"},
//			},
//			VolumeMounts: []apiv1.VolumeMount{
//				{Name: "log", MountPath: "/var/log/nginx"},
//				// TODO js path
//				{Name: "nginx-stack-config", MountPath: "/etc/nginx/templates/nginx.conf.template", SubPath: "nginx.conf"},
//				{Name: "nginx-stack-config", MountPath: "/etc/nginx/index.js", SubPath: "index.js"},
//			},
//		},
//	}
//	for _, function := range functions {
//		//function.Port = strconv.Itoa(currentPort)
//		//container = append(container, apiv1.Container{
//		//	Name:  function.Name,
//		//	Image: function.Image,
//		//	Env: []apiv1.EnvVar{
//		//		{
//		//			Name:  "PORT",
//		//			Value: function.Port,
//		//		},
//		//	},
//		//})
//		//currentPort += currentPort
//
//	}
//	bytes, _ := json.Marshal(functions)
//	container[0].Env = append(container[0].Env, apiv1.EnvVar{Name: "FUNCTION_DATA", Value: base64.StdEncoding.EncodeToString(bytes)})
//
//	return container
//}
