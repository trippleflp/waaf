package handler

import (
	"fmt"
	"function-uploader/internal/oci"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
)

func UploadHandler(c *fiber.Ctx) error {
	functionGroupId := c.Params("functionGroup")
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

	location, err := oci.Builder().
		SetFile(wasmPath).
		SetRegistryUrl("http://kind-registry:5000").
		SetTag("1.0.0").
		SetName(fmt.Sprintf("%s/%s", functionGroupId, functionName)).
		Build()
	if err != nil {
		log.Err(err)
		return err
	}

	return c.SendString(location)

}
