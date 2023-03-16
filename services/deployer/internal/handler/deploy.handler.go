package handler

import (
	"deployer/internal/deployment"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
)

type DeployHandlerBody struct {
	Functions         []string `json:"functions"`
	FunctionGroupName string   `json:"functionGroupName""`
}

const registryUrl = "localhost:5001"

func DeployHandler(c *fiber.Ctx) error {
	body := new(DeployHandlerBody)
	err := c.BodyParser(body)
	if err != nil {
		log.Debug().Err(err).Str("body", string(c.Body())).Msg("Body parsing did not work")
		return fiber.NewError(fiber.StatusBadRequest, "Body can't be parsed")
	}

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

	deploymentMangerBuilder := deployment.
		Builder(clientset).
		SetContext(c.UserContext()).
		SetFunctionGroupName(body.FunctionGroupName)

	for _, fn := range body.Functions {
		deploymentMangerBuilder.SetFunctions(strings.Split(strings.Split(fn, "/")[1], ":")[0], fmt.Sprintf("%s/%s", registryUrl, fn))
	}
	deploymentManger, err := deploymentMangerBuilder.Build()

	if err != nil {
		log.Err(err)
		return err
	}

	err = deploymentManger.DeployAll()
	if err != nil {
		log.Err(err)
		return err
	}

	res := fmt.Sprintf("Created deployment of functiongroup: %s.\n", body.FunctionGroupName)
	fmt.Println(res)
	return c.SendString(res)
}
