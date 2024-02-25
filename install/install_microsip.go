package install

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/execute"
	"github.com/agnerft/ListRamais/util"
)

func InstallMicrosip(cliente *domain.Cliente, ramal domain.Ramal) error {

	var err error
	var destDeleleteMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "Uninstall.exe")
	var pathMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP")
	var destDownMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "MicroSIP-3.21.3.exe")
	var destFileConfigMicrosip = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "microsip.ini")
	var url = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"

	err = util.Executable(destDeleleteMicroSIP)
	if err != nil {
		fmt.Printf("Erro ou executar o Desinstalador.")
	}

	err = execute.DownloadGeneric(url, destDownMicroSIP)
	if err != nil {

		return err
	}
	err = util.Executable(destDownMicroSIP)
	if err != nil {
		log.Printf("Erro ao instalar o %s", destDownMicroSIP)
	}
	if _, err = os.Stat(pathMicroSIP); err == nil {

		fmt.Printf("Já existe os arquivos na pasta %s", pathMicroSIP)

	} else if os.IsNotExist(err) {
		fmt.Println("o caminho não existe")

		// BAIXAR O MICROSIP
		err = execute.DownloadGeneric(url, destDownMicroSIP)
		if err != nil {
			log.Fatal("Erro ao baixar o Arquivo.", err)
		}

		err = util.Executable(destDownMicroSIP)
		if err != nil {
			fmt.Printf("Erro ou executar o Instalador.")
		}

		fmt.Print("Aguardando")

	} else {
		// Algum erro ocorreu ao verificar o caminho
		fmt.Printf("Erro ao verificar o caminho: %v\n", err)

	}

	err = util.AdicionarConfiguracao(destFileConfigMicrosip)
	if err != nil {
		log.Fatal("Erro ao Adicionar a Configuração. \n", err)
	}
	fmt.Println("Passou aqui?")
	// EDIÇÃO DO ARQUIVO

	replace := map[string]string{
		"accountId=0":         "accountId=1",
		"videoBitrate=0":      "videoBitrate=256",
		"recordingPath=":      fmt.Sprintf("recordingPath=%s", filepath.Join(util.UserCurrent().HomeDir, "Desktop")),
		"recordingFormat=":    "recordingFormat=mp3",
		"autoAnswer=button":   "autoAnswer=all",
		"denyIncoming=button": "denyIncoming=",
		"label=":              fmt.Sprintf("label=%s", ramal.Sip),
		"server=":             fmt.Sprintf("server=%s", cliente.Link_sip),
		"proxy=":              fmt.Sprintf("proxy=%s", cliente.Link_sip),
		"domain=":             fmt.Sprintf("domain=%s", cliente.Link_sip),
		"username=":           fmt.Sprintf("username=%s", ramal.Sip),
		"password=":           fmt.Sprintf("password=%s@abc", ramal.Sip),
		"authID=":             fmt.Sprintf("authID=%s", ramal.Sip),
	}

	for key, value := range replace {
		err := util.ReplaceLineOfFile(destFileConfigMicrosip, key, value)
		if err != nil {
			log.Printf("Erro ao modificar %s. %s \n", key, err)
		}
	}
	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "accountId=0", "accountId=1")
	// if err != nil {
	// 	log.Printf("Erro para modificar o AccountId. %s \n", err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "videoBitrate=0", "videoBitrate=256")
	// if err != nil {
	// 	log.Printf("Erro para modificar o videoBitrate. %s \n", err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "recordingPath=", fmt.Sprintf("%s%s", "recordingPath=", filepath.Join(util.UserCurrent().HomeDir, "Desktop")))
	// if err != nil {
	// 	log.Printf("Erro para modificar o recordingPath. %s \n", err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "recordingFormat=", "recordingFormat=mp3")
	// if err != nil {
	// 	log.Printf("Erro para modificar o recordingFormat. %s \n", err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "autoAnswer=button", "autoAnswer=all")
	// if err != nil {
	// 	log.Printf("Erro para modificar o autoAnswer. %s \n", err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "denyIncoming=button", "denyIncoming=")
	// if err != nil {
	// 	log.Printf("Erro para modificar o denyIncoming. %s \n", err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "label=", fmt.Sprintf("%s%s", "label=", ramal.Sip))
	// if err != nil {
	// 	log.Printf("Erro para modificar o Sip no Label. %s \n", err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "server=", fmt.Sprintf("%s%s", "server=", cliente.Link_sip))
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. \n %s", cliente.Cliente, err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "proxy=", fmt.Sprintf("%s%s", "proxy=", cliente.Link_sip))
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. \n %s", cliente.Cliente, err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "domain=", fmt.Sprintf("%s%s", "domain=", cliente.Link_sip))
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. %s", cliente.Cliente, err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "username=", fmt.Sprintf("%s%s", "username=", ramal.Sip))
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. %s", cliente.Cliente, err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "password=", fmt.Sprintf("%s%s%s", "password=", ramal.Sip, "@abc"))
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. %s", cliente.Cliente, err)
	// }

	// err = util.ReplaceLineOfFile(destFileConfigMicrosip, "authID=", fmt.Sprintf("%s%s", "authID=", ramal.Sip))
	// if err != nil {
	// 	log.Printf("Erro para setar o link do cliente %s. %s", cliente.Cliente, err)
	// }

	return nil
}

func ReadIntheFile(cliente *domain.Cliente, ramal domain.Ramal) {

}
