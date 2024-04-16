package handler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

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

// func HandleCallWhatsApp(c *fiber.Ctx) error {
// 	cnpj := c.Params("cnpj")
// 	ramal := c.Params("ramal")
// 	urlWhatsapp := fmt.Sprintf("%s%s%s%s%s", `https://api.whatsapp.com/send/?phone=555130655144&text=cnpj%3D`, cnpj, `%3Bramal%3D`, ramal, `&type=phone_number&app_absent=0`)
// 	// `https://api.whatsapp.com/send/?phone=555130655144&text=cnpj%3D20905507000100%3Bramal%3D7801&type=phone_number&app_absent=0`
// 	err := util.OpenBrowser(urlWhatsapp)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(c.Status(200))
// }

func HandleClient(c *fiber.Ctx) error {
	var response = map[string]interface{}{}
	cnpj := c.Params("cnpj")
	fmt.Println(cnpj)
	// fmt.Println(response)
	newCliente, err := getClient(cnpj)
	if err != nil {
		return err

	}

	newRamais, err := getRamais(cnpj)
	if err != nil {
		return err
	}

	response = map[string]interface{}{
		"json1": newCliente,
		"json2": newRamais,
	}

	fmt.Println(response)

	return c.JSON(response)

}

func getRamais(cnpj string) (domain.RamalSolo, error) {

	newClient, err := getClient(cnpj)
	if err != nil {
		return domain.RamalSolo{}, err
	}

	if &cliente != nil {
		newRamais, err := svc.PostRamais(fmt.Sprintf("%s/%s", newClient.Link, "asterisk_exec"))
		if err != nil {
			return domain.RamalSolo{}, err
		}

		return newRamais, nil
	}

	return domain.RamalSolo{}, nil
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

	newRamais, err := svc.PostRamais(fmt.Sprintf("%s/%s", newCliente.Link, "asterisk_exec"))
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
	ramalParamInt, _ := strconv.Atoi(ramalParam)
	acc := c.Params("acc")

	newCliente, err := getClient(cnpj)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newRamais, err := svc.PostRamais(fmt.Sprintf("%s/%s", newCliente.Link, "asterisk_exec"))
	if err != nil {
		return err
	}

	var ramalAtual domain.Ramal

	for _, ramal := range newRamais.Ramais {

		if ramal.Sip == ramalParamInt {
			fmt.Println(ramal)
			// fmt.Println(ramalAtual.Ramais)
			// ramalAtual.Ramais = append(ramalAtual.Ramais, ramal)

			ramalAtual = ramal

			break
		}

	}

	r, err := install.InstallMicrosip(newCliente, ramalAtual, fmt.Sprintf("%s%s", "Account", acc))
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
	fmt.Println(r)
	fmt.Println(response)

	return c.JSON(response)
}

func HandlerUninstall(c *fiber.Ctx) error {

	var pathMicroSIP = filepath.Join(util.UserCurrent().HomeDir,
		"AppData",
		"Local",
		"MicroSIP",
		"microsip.exe")

	pid, err := util.GetPIDbyName(filepath.Base(pathMicroSIP))
	if err != nil {
		return err
	}
	fmt.Println(pid)

	if pid != 0 {
		err = util.TaskkillExecute(pid)
		if err != nil {
			return err
		}
	}

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

	dir, err := os.Open(confiMicroSIP)
	if err != nil {
		fmt.Println("Deu erro para ler")
	}

	defer dir.Close()

	files, err := dir.Readdirnames(-1)
	if err != nil {

		fmt.Println("Erro ao ler os nomes dos arquivos.")
	}

	if len(files) == 0 {
		fmt.Println("Pasta vazia")
	} else {
		// fmt.Println("Arquivos:")
		for _, nameFile := range files {
			err := os.Remove(fmt.Sprintf("%s/%s", confiMicroSIP, nameFile))
			if err != nil {
				fmt.Println("Deu erro pra excluir")
			}
			fmt.Printf("Removido o arquivo: %s", nameFile)
		}
	}

	return c.JSON(response)

}

func HandlerSaveFile(c *fiber.Ctx) error {
	cnpj := c.Params("cnpj")

	newClient, err := getClient(cnpj)
	if err != nil {
		return err
	}

	ramais, err := getRamais(cnpj)
	if err != nil {
		return err
	}

	err = util.FileInfos(*newClient, ramais)
	if err != nil {
		return err
	}

	response := map[string]interface{}{
		"cliente": newClient.Cliente,
		"ramais":  ramais.Ramais,
	}

	return c.JSON(response)

}
