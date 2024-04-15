package router

import (
	_ "embed"
	"fmt"

	"github.com/agnerft/ListRamais/handler"
	"github.com/gofiber/fiber/v2"
)

func InitRouter() error {
	app := fiber.New(fiber.Config{})

	app.Get("/:cnpj", handler.HandleClient)
	app.Get("/:cnpj/all", handler.HandleRamais)
	app.Get("/:cnpj/:ramal", handler.HandleCallWhatsApp)
	app.Get("/:cnpj/:ramal/install/:acc", handler.HandlerInstall)
	app.Get("/:cnpj/uninstall", handler.HandlerUninstall)
	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
