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
