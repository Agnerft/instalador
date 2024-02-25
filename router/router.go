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
	app.Get("/:cnpj/ramais", handler.HandleRamais)
	app.Get("/:cnpj/:ramal/install", handler.HandlerInstall)
	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
