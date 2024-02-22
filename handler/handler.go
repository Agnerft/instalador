package handler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/execute"
	"github.com/agnerft/ListRamais/services"
	"github.com/agnerft/ListRamais/util"
	"github.com/gofiber/fiber/v2"
)

var (
	cliente *domain.Cliente

	url                  = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	destDeleleteMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "Uninstall.exe")
	// destRunningMicroSIP    = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "microsip.exe")
	destDownMicroSIP       = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "MicroSIP-3.21.3.exe")
	destFileConfigMicrosip = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "microsip.ini")
	ramalAtual             string
	processName            = "microsip.exe"
	svc                    = services.NewServiceCliente()
	ramaisDoCliente        []string
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
	fmt.Println(cnpj)
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

func InstallMicrosip(cliente domain.Cliente) {

	var err error

	if _, err = os.Stat(destDeleleteMicroSIP); err == nil {
		// Se o caminho existe, execute algo

		err = util.ExecuteUnistall(destDeleleteMicroSIP)
		if err != nil {
			fmt.Printf("Erro ou executar o Desinstalador.")
		}

		err = execute.DownloadGeneric(url, destDownMicroSIP)
		if err != nil {
			log.Fatal("Erro ao baixar o Arquivo.", err)
		}

		err = util.ExecuteInstall(destDownMicroSIP)
		if err != nil {
			log.Printf("Erro ao instalar o %s", destDownMicroSIP)
		}

	} else if os.IsNotExist(err) {
		fmt.Println("o caminho não existe")
		// Se o caminho não existe, faça algo diferente
		// BAIXAR O MICROSIP
		err = execute.DownloadGeneric(url, destDownMicroSIP)
		if err != nil {
			log.Fatal("Erro ao baixar o Arquivo.", err)
		}

		err = util.ExecuteInstall(destDownMicroSIP)
		if err != nil {
			fmt.Printf("Erro ou executar o Instalador.")
		}

		fmt.Print("Aguardando")

	} else {
		// Algum erro ocorreu ao verificar o caminho
		fmt.Printf("Erro ao verificar o caminho: %v\n", err)
		// Adicione aqui o código para lidar com o erro, se necessário
	}
	// http://localhost:8080/20905507000100/2365/install
	// :cnpj/:ramal/install
	fmt.Printf("Chamando configuração")

	fmt.Println("teste")
}

// func HandleFileConfig(c *gin.Context) {

// 	fmt.Println("Ta clicando aqui?")

// 	// fmt.Println(RamalSelecionado)

// 	err := util.AdicionarConfiguracao(destFileConfigMicrosip)
// 	if err != nil {
// 		log.Fatal("Erro ao Adicionar a Configuração. \n", err)

// 	}

// 	fmt.Println(string(Cliente.Cliente))
// 	// EDIÇÃO DO ARQUIVO

// 	err = util.ReadFile(destFileConfigMicrosip, "accountId=0", "accountId=1", 1)
// 	if err != nil {
// 		log.Fatal("Erro para modificar o AccountId. \n", err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "videoBitrate=0", "videoBitrate=256", 23)
// 	if err != nil {
// 		log.Fatal("Erro para modificar o videoBitrate. \n", err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "recordingPath=", "recordingPath="+filepath.Join(util.UserCurrent().HomeDir, "Desktop"), 32)
// 	if err != nil {
// 		log.Fatal("Erro para modificar o recordingPath. \n", err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "recordingFormat=", "recordingFormat=mp3", 33)
// 	if err != nil {
// 		log.Fatal("Erro para modificar o recordingFormat. \n", err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "autoAnswer=button", "autoAnswer=all", 37)
// 	if err != nil {
// 		log.Fatal("Erro para modificar o autoAnswer. \n", err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "denyIncoming=button", "denyIncoming=", 43)
// 	if err != nil {
// 		log.Fatal("Erro para modificar o denyIncoming. \n", err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "label=", "label="+ramalAtual, 106)
// 	if err != nil {
// 		log.Fatal("Erro para modificar o Sip no Label. \n", err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "server=", "server="+string(Cliente.Link_sip), 107)
// 	if err != nil {
// 		log.Fatalf("Erro para setar o link do cliente %s. \n %s", string(Cliente.Cliente), err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "proxy=", "proxy="+string(Cliente.Link_sip), 108)
// 	if err != nil {
// 		log.Fatalf("Erro para setar o link do cliente %s. \n %s", string(Cliente.Cliente), err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "domain=", "domain="+string(Cliente.Link_sip), 109)
// 	if err != nil {
// 		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "username=", "username="+ramalAtual, 110)
// 	if err != nil {
// 		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "password=", "password="+ramalAtual+"@abc", 111)
// 	if err != nil {
// 		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
// 	}

// 	err = util.ReadFile(destFileConfigMicrosip, "authID=", "authID="+ramalAtual, 112)
// 	if err != nil {
// 		log.Printf("Erro para setar o link do cliente %s. %s", string(Cliente.Cliente), err)
// 	}

// }

// func HandleTeste(c *gin.Context) {
// 	cmd := exec.Command("ps", "-p", fmt.Sprint(890))
// 	err := cmd.Run()

// 	if err != nil {
// 		fmt.Printf("Deu erro: %s", err)
// 	}
// }
