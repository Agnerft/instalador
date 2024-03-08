package handler

import (
	"fmt"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/install"
	"github.com/agnerft/ListRamais/services"
	"github.com/gofiber/fiber/v2"
)

func HandlePingPong(c *fiber.Ctx) error {
	resp := map[string]string{
		"resp": "pong",
	}
	return c.JSON(resp)
}

func HandleClient(c *fiber.Ctx) error {

	cnpj := c.Params("cnpj")
	fmt.Println(cnpj)
	newCliente, err := getClient(cnpj)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)

	}

	newRamais, err := getRamais(cnpj)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	response := map[string]interface{}{
		"json1": newCliente,
		"json2": newRamais,
	}

	fmt.Println(response)

	return c.JSON(response)

}

func getRamais(cnpj string) (domain.RamaisRegistrados, error) {

	svc := services.NewServiceCliente()
	var cliente *domain.Cliente

	newClient, err := getClient(cnpj)
	if err != nil {
		return domain.RamaisRegistrados{}, err
	}

	if cliente == nil {
		newRamais, err := svc.RequestJsonRamal(newClient.Link)
		if err != nil {
			return domain.RamaisRegistrados{}, err
		}

		return newRamais, nil
	}

	return domain.RamaisRegistrados{}, nil
}

func getClient(cnpj string) (*domain.Cliente, error) {

	svc := services.NewServiceCliente()
	var cliente *domain.Cliente

	if cliente == nil || cliente.Documento != cnpj {

		cliente, err := svc.RequestJsonCliente(cnpj)
		if err != nil {
			return nil, err

		} else {
			return &cliente, nil
		}
	} else {
		return cliente, nil
	}

}

func HandleRamais(c *fiber.Ctx) error {
	svc := services.NewServiceCliente()
	cnpj := c.Params("cnpj")

	newCliente, err := getClient(cnpj)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)

	}

	newRamais, err := svc.RequestJsonRamal(newCliente.Link)
	if err != nil {
		return err
	}

	return c.JSON(newRamais)

}

func HandlerInstall(c *fiber.Ctx) error {
	svc := services.NewServiceCliente()
	cnpj := c.Params("cnpj")
	ramalParam := c.Params("ramal")
	acc := c.Params("acc")

	newCliente, err := getClient(cnpj)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newRamais, err := svc.RequestJsonRamal(newCliente.Link)
	if err != nil {
		return err
	}

	var ramalAtual domain.Ramal

	for _, ramal := range newRamais.RamaisRegistrados {
		if ramal.Sip == ramalParam {
			fmt.Println(ramal.Sip)
			ramalAtual = ramal
			break
		}

	}

	str, err := install.InstallMicrosip(newCliente, ramalAtual, fmt.Sprintf("%s%s", "Account", acc))
	if err != nil {
		return err
	}

	return c.JSON(str)
}
