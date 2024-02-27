package install

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/execute"
	"github.com/agnerft/ListRamais/util"
)

func InstallMicrosip(cliente *domain.Cliente, ramal domain.Ramal, account string) (string, error) {

	var err error
	var destDeleleteMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "Uninstall.exe")
	var pathMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "microsip.exe")
	var destDownMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "MicroSIP-3.21.3.exe")
	var destFileConfigMicrosip = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "MicroSIP.ini")
	var url = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"

	fmt.Println(account)
	err = util.Executable(destDeleleteMicroSIP)
	if err != nil {
		fmt.Printf("Erro ou executar o Desinstalador no caminho: %s. \n", destDeleleteMicroSIP)
	}

	err = execute.DownloadGeneric(url, destDownMicroSIP)
	if err != nil {

		return "", err
	}

	err = util.Executable(destDownMicroSIP)
	if err != nil {
		log.Printf("Erro ao executar o instalador no caminho: %s. \n", destDownMicroSIP)
	}

	i, err := util.GetPIDbyName(filepath.Base(pathMicroSIP))
	if err != nil {
		return "", err
	}

	err = util.TaskkillExecute(i)
	if err != nil {
		return "", err
	}

	ini := util.NewIniFile(destFileConfigMicrosip)

	err = ini.Readini()
	if err != nil {
		fmt.Println(err)
	}

	cfg := domain.NewConfig()

	if !ini.ExistsVauleFromSectionAndKey("Settings", "videoBitrate", "256") {

		mpConfigSettings := make(map[string]string, 0)
		mpConfigSettings["videoBitrate"] = "256"
		mpConfigSettings["recordingPath"] = filepath.Join(util.UserCurrent().HomeDir, "Desktop")
		mpConfigSettings["recordingFormat"] = "mp3"
		mpConfigSettings["autoAnswer"] = "all"
		mpConfigSettings["denyIncoming"] = ""
		ini.UpdateBatchSection("Settings", mpConfigSettings)

	}

	cfg.Label = ramal.Sip
	cfg.Server = cliente.Link_sip
	cfg.Proxy = cliente.Link_sip
	cfg.Domain = cliente.Link_sip
	cfg.Username = ramal.Sip
	cfg.Password = fmt.Sprintf("%s%s", ramal.Sip, "@abc")
	cfg.AuthID = ramal.Sip

	fmt.Println(*cfg)

	ini.AddSectionAccount(account, *cfg)

	err = ini.WriteIni()
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("Instalado MicroSIP com o ramal %s", ramal.Sip), nil
}
