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
	//defer os.RemoveAll(temp)

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

	err = os.WriteFile(wasmPath, bytes, 0644)
	if err != nil {
		return err
	}

	err = oci.Builder().
		SetFile(wasmPath).
		SetRegistryUrl("http://kind-registry:5000").
		SetTag("1.0.0").
		SetName("test/wasm_exec").
		Build()
	if err != nil {
		log.Err(err)
		return err
	}

	return c.SendString(fmt.Sprintf("%s:%s", "test/wasm_exec", "1.0.0"))

}
