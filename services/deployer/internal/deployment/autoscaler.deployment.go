package deployment

import (
	v1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getAutoScaler() {

	v1.HorizontalPodAutoscaler{
		ObjectMeta: metav1.ObjectMeta{
			Name: "function name",
		},
		Spec: v1.HorizontalPodAutoscalerSpec{
			ScaleTargetRef: v1.CrossVersionObjectReference{
				Kind: "Deployment",
				APIVersion: "apps/v1",
				Name: "function name",
			},
			MinReplicas: 0
		},
	}
}