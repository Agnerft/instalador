package install

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/agnerft/ListRamais/domain"
	"github.com/agnerft/ListRamais/execute"
	"github.com/agnerft/ListRamais/util"
)

func InstallMicrosip(cliente *domain.Cliente, ramal domain.ObjetoGvc, account string) (string, error) {
	duration := 5 * time.Second
	var err error
	// var destDeleleteMicroSIP = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Local", "MicroSIP", "Uninstall.exe")
	var pathMicroSIP = filepath.Join(util.UserCurrent().HomeDir,
		"AppData",
		"Local",
		"MicroSIP",
		"microsip.exe")
	var destDownMicroSIP = filepath.Join(util.UserCurrent().HomeDir,
		"AppData",
		"Local",
		"MicroSIP",
		"MicroSIP-3.21.3.exe")
	var destFileConfigMicrosip = filepath.Join(util.UserCurrent().HomeDir,
		"AppData",
		"Roaming",
		"MicroSIP",
		"MicroSIP.ini")
	var url = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"

	fmt.Println(account)

	confiMicroSIP := filepath.Join(util.UserCurrent().HomeDir,
		"AppData",
		"Roaming",
		"MicroSIP")

	err = util.CleanFiles(confiMicroSIP)
	if err != nil {
		return "", err
	}

	err = execute.DownloadGeneric(url, destDownMicroSIP)
	if err != nil {
		return "", err
	}

	err = util.Executable(destDownMicroSIP)
	if err != nil {
		log.Printf("Erro ao executar o instalador no caminho: %s. \n", destDownMicroSIP)
	}

	time.Sleep(duration)

	i, err := util.GetPIDbyName(filepath.Base(pathMicroSIP))
	if err != nil {
		return "", err
	}

	fmt.Println(i)
	// err = util.OpenMicroSIP(pathMicroSIP)
	// if err != nil {
	// 	return "", err
	// }

	err = util.TaskkillExecute(i)
	if err != nil {
		return "", err
	}

	if !util.FileIsExist(destFileConfigMicrosip) {

		err = util.OpenMicroSIP(pathMicroSIP)
		if err != nil {
			return "", err
		}

		time.Sleep(duration)

		err = util.TaskkillExecute(i)
		if err != nil {
			return "", err
		}

	}

	ini := util.NewIniFile(destFileConfigMicrosip)

	err = ini.Readini()
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(duration)

	mpConfigSettings := make(map[string]string, 0)
	set := "Settings"

	// fmt.Println(ini.ExistsKey(set, rec))

	mpConfigSettings["recordingPath"] = filepath.Join(util.UserCurrent().HomeDir, "Desktop")
	mpConfigSettings["recordingFormat"] = "mp3"
	mpConfigSettings["autoAnswer"] = "all"
	mpConfigSettings["denyIncoming"] = ""
	fmt.Println(ini.UpdateBatchSection(set, mpConfigSettings))

	cfg := domain.NewConfig()
	// ramalString := strconv.Itoa(ramal.NameUsername)

	cfg.Label = ramal.NameUsername
	cfg.Server = cliente.Link_sip
	cfg.Proxy = cliente.Link_sip
	cfg.Domain = cliente.Link_sip
	cfg.Username = ramal.NameUsername
	cfg.Password = fmt.Sprintf("%s%s", ramal.NameUsername, "@abc")
	cfg.AuthID = ramal.NameUsername

	// fmt.Println(*cfg)

	ini.AddSectionAccount(account, *cfg)

	err = ini.WriteIni()
	if err != nil {
		fmt.Println(err)
	}

	// err = util.OpenMicroSIP(pathMicroSIP)
	// if err != nil {
	// 	return "", err
	// }

	// i2, err := util.GetPIDbyName(filepath.Base(pathMicroSIP))
	// if err != nil {
	// 	return "", err
	// }

	// fmt.Println(i2)
	// // err = util.OpenMicroSIP(pathMicroSIP)
	// // if err != nil {
	// // 	return "", err
	// // }

	// err = util.TaskkillExecute(i2)
	// if err != nil {
	// 	return "", err
	// }

	return fmt.Sprintf("Instalado MicroSIP com o ramal %s", ramal.NameUsername), nil
}
