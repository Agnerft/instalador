package handler

import (
	"fmt"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/install"
	"github.com/agnerft/ListRamais/services"
	"github.com/gofiber/fiber/v2"
)

var (
	cliente *domain.Cliente

	// url                  = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	// destDeleleteMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "Uninstall.exe")
	// destRunningMicroSIP    = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "microsip.exe")
	// destDownMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "MicroSIP-3.21.3.exe")
	// destFileConfigMicrosip = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "microsip.ini")
	// ramalAtual             string
	// processName     = "microsip.exe"
	svc             = services.NewServiceCliente()
	ramaisDoCliente []string
)

func HandleClient(c *fiber.Ctx) error {

	cnpj := c.Params("cnpj")
	fmt.Println(cnpj)
	newCliente, err := getClient(cnpj)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)

	}

	return c.JSON(newCliente)

}

func getClient(cnpj string) (*domain.Cliente, error) {

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
	cnpj := c.Params("cnpj")
	ramalParam := c.Params("ramal")

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

	err = install.InstallMicrosip(newCliente, ramalAtual, "Account1")
	if err != nil {
		return err
	}

	return c.JSON(newCliente)
}
