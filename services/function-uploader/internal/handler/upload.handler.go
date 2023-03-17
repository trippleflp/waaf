package handler

import (
	"fmt"
	"function-uploader/internal/oci"
	"github.com/gofiber/fiber/v2"
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
)

type AddFunctionBody struct {
	FunctionGroupName string `json:"functionGroupName"`
	FunctionTag       string `json:"functionTag"`
}

func UploadHandler(c *fiber.Ctx) error {
	functionGroupName := c.Params("functionGroup")
	functionName := c.Params("functionName")
	file, err := c.FormFile("fileUpload")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	temp, err := os.MkdirTemp("", "wasm-upload-*")
	if err != nil {
		return err
	}

	fmt.Println("open file")
	reader, err := file.Open()
	if err != nil {
		log.Err(err)
		return err
	}
	defer reader.Close()

	fmt.Println("read file")
	bytes, err := io.ReadAll(reader)
	if err != nil {
		log.Err(err)
		return err
	}
	wasmFileName := "wasm_exec.wasm"
	wasmPath := filepath.Join(temp, wasmFileName)
	fmt.Println("write file to ", wasmPath)

	err = os.WriteFile(wasmPath, bytes, 0777)
	if err != nil {
		log.Err(err)
		return err
	}

	image, err := oci.Builder().
		SetFile(wasmPath).
		SetRegistryUrl("http://kind-registry:5000").
		SetTag("latest").
		SetName(fmt.Sprintf("%s/%s", functionGroupName, functionName)).
		Build()
	if err != nil {
		log.Err(err)
		return err
	}

	url, exist := os.LookupEnv("FUNCTIONGROUP_URL")
	if !exist {
		return fmt.Errorf("function group url not set")
	}

	body := &AddFunctionBody{
		FunctionGroupName: functionGroupName,
		FunctionTag:       image,
	}

	response, err := req.R().SetBody(body).Post(fmt.Sprintf("%s/groups/%s/addFunction", url, functionGroupName))
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("function upload failed")
	}

	return c.SendString("function upload was successfully")

}
