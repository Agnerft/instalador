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

func InstallMicrosip(cliente *domain.Cliente, ramal domain.Ramal, account string) error {

	var err error
	// var destDeleleteMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "Uninstall.exe")
	var pathMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP")
	var destDownMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "MicroSIP-3.21.3.exe")
	var destFileConfigMicrosip = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "MicroSIP.ini")
	var url = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"

	// err = util.Executable(destDeleleteMicroSIP)
	// if err != nil {
	// 	fmt.Printf("Erro ou executar o Desinstalador.")
	// }

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

	// err = util.AdicionarConfiguracao(destFileConfigMicrosip)
	// if err != nil {
	// 	log.Fatal("Erro ao Adicionar a Configuração. \n", err)
	// }
	fmt.Println("Passou aqui?")
	// EDIÇÃO DO ARQUIVO
	ini := util.NewIniFile(destFileConfigMicrosip)

	err = ini.Readini()
	if err != nil {
		fmt.Println(err)
	}

	cfg := domain.NewConfig()

	mpConfigSettings := make(map[string]string, 0)
	mpConfigSettings["accountId"] = account
	mpConfigSettings["videoBitrate"] = "256"
	mpConfigSettings["recordingPath"] = filepath.Join(util.UserCurrent().HomeDir, "Desktop")
	mpConfigSettings["recordingFormat"] = "mp3"
	mpConfigSettings["autoAnswer"] = "all"
	mpConfigSettings["denyIncoming"] = ""
	ini.UpdateBatchSection("Settings", mpConfigSettings)

	cfg.Label = ramal.Sip
	cfg.Server = cliente.Link_sip
	cfg.Proxy = cliente.Link_sip
	cfg.Domain = cliente.Link_sip
	cfg.Username = ramal.Sip
	cfg.Password = fmt.Sprintf("%s%s", ramal.Sip, "@abc")
	cfg.AuthID = ramal.Sip

	fmt.Println(cfg)

	ini.AddSectionAccount(account, *cfg)

	err = ini.WriteIni()
	if err != nil {
		fmt.Println(err)
	}

	// replace := map[string]string{
	// 	"accountId":           "accountId=1",
	// 	"videoBitrate":        "videoBitrate=256",
	// 	"recordingPath":       fmt.Sprintf("recordingPath=%s", filepath.Join(util.UserCurrent().HomeDir, "Desktop")),
	// 	"recordingFormat":     "recordingFormat=mp3",
	// 	"autoAnswer=button":   "autoAnswer=all",
	// 	"denyIncoming=button": "denyIncoming=",
	// 	"label":               fmt.Sprintf("label=%s", ramal.Sip),
	// 	"server":              fmt.Sprintf("server=%s", cliente.Link_sip),
	// 	"proxy":               fmt.Sprintf("proxy=%s", cliente.Link_sip),
	// 	"domain":              fmt.Sprintf("domain=%s", cliente.Link_sip),
	// 	"username":            fmt.Sprintf("username=%s", ramal.Sip),
	// 	"password":            fmt.Sprintf("password=%s@abc", ramal.Sip),
	// 	"authID":              fmt.Sprintf("authID=%s", ramal.Sip),
	// }

	return nil
}
