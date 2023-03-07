package handler

import (
	"deployer/internal/deployment"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func DeployHandler(c *fiber.Ctx) error {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	deploymentManger, err := deployment.
		Builder(clientset).
		SetContext(c.UserContext()).
		SetFunctionGroupName("waaf").
		SetFunctions("hello", "localhost:5001/hello-app:1.0").
		Build()

	if err != nil {
		return err
	}

	result, err := deploymentManger.DeployNginx()
	if err != nil {
		return err
	}

	res := fmt.Sprintf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	fmt.Println(res)
	return c.SendString(res)
	//
	//for {
	//	// get pods in all the namespaces by omitting namespace
	//	// Or specify namespace to get pods in particular namespace
	//	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//
	//	// Examples for error handling:
	//	// - Use helper functions e.g. errors.IsNotFound()
	//	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	//	_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
	//	if errors.IsNotFound(err) {
	//		fmt.Printf("Pod example-xxxxx not found in default namespace\n")
	//	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	//		fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
	//	} else if err != nil {
	//		panic(err.Error())
	//	} else {
	//		fmt.Printf("Found example-xxxxx pod in default namespace\n")
	//	}
	//
	//	time.Sleep(10 * time.Second)
	//}
	//return nil
}
