package handler

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/install"
	"github.com/agnerft/ListRamais/services"
	"github.com/agnerft/ListRamais/util"
	"github.com/gofiber/fiber/v2"
)

var (
	cliente domain.Cliente
	svc     services.ServiceRequest
)

func init() {
	svc = *services.NewServiceCliente()
}

func HandlePingPong(c *fiber.Ctx) error {
	resp := map[string]string{
		"resp": "pong",
	}
	return c.JSON(resp)
}

func HandleClient(c *fiber.Ctx) error {
	var response = map[string]interface{}{}
	cnpj := c.Params("cnpj")
	fmt.Println(cnpj)
	// fmt.Println(response)
	newCliente, err := getClient(cnpj)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)

	}

	newRamais, err := getRamais(cnpj)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	response = map[string]interface{}{
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

	if &cliente != nil {
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

	if &cliente != nil || cliente.Documento != cnpj {

		cliente, err := svc.RequestJsonCliente(cnpj)
		if err != nil {
			return nil, err

		} else {
			return &cliente, nil
		}
	} else {
		return &cliente, nil
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

	var response = map[string]any{}

	confiMicroSIP := filepath.Join(util.UserCurrent().HomeDir,
		"AppData",
		"Local",
		"MicroSIP",
	)

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

	_, err = install.InstallMicrosip(newCliente, ramalAtual, fmt.Sprintf("%s%s", "Account", acc))
	if err != nil {
		return err
	}

	exist := util.FileIsExist(confiMicroSIP)

	response = map[string]interface{}{
		"cliente": newCliente.Cliente,
		"doc":     newCliente.Documento,
		"ramal":   ramalAtual.Sip,
		"status":  exist,
	}

	return c.JSON(response)
}

func HandlerUninstall(c *fiber.Ctx) error {

	var response = map[string]bool{}

	destDeleleteMicroSIP := filepath.Join(util.UserCurrent().HomeDir,
		"AppData",
		"Local",
		"MicroSIP",
		"Uninstall.exe")

	confiMicroSIP := filepath.Join(util.UserCurrent().HomeDir,
		"AppData",
		"Roaming",
		"MicroSIP")

	if util.FileIsExist(destDeleleteMicroSIP) {
		err := util.Executable(destDeleleteMicroSIP)
		if err != nil {
			log.Fatal("Erro ao executar o Desinstalador.")

		}

	}

	if util.FileIsExist(confiMicroSIP) {
		err := util.RemovePath(confiMicroSIP)
		if err != nil {
			fmt.Printf("Erro ao remover a Pasta %s", confiMicroSIP)

		}
	}

	exist := util.FileIsExist(destDeleleteMicroSIP)

	response = map[string]bool{
		"removido": exist,
	}

	fmt.Println(response)

	return c.JSON(response)

}
