package util

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

// var UsrCurr *user.User

func ExecuteUnistall(filePath string) error {
	// C:\Users\USER\AppData\Local\MicroSIP\Uninstall.exe

	// filePath := filepath.Join(UserCurrent().HomeDir,
	// 	"AppData",
	// 	"Local",
	// 	"MicroSIP",
	// 	"Uninstall.exe")

	cmd := exec.Command(filePath, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o desinstalador: %s \n", err)
		return err
	}

	fmt.Println("Removido")

	return nil
}

func ExecuteInstall(filePath string) error {
	cmd := exec.Command(filePath, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o desinstalador: %s \n", err)
		return err
	}

	fmt.Println("Instalado")

	return nil
}

func UserCurrent() user.User {
	// Obter o diretório do usuário
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Erro ao obter o diretório do usuário: \n", err)
	}

	return *usr
}

func OpenBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		fmt.Printf("Não foi possível detectar o sistema operacional para abrir o navegador automaticamente.")
		return
	}

	if err != nil {
		fmt.Printf("Erro ao abrir o navegador: %v\n", err)
	}
}

func TaskkillExecute(pid int) error {
	cmd := exec.Command("taskkill", "/pid", strconv.Itoa(pid)) //TASKKILL /IM microsip.exe
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o TASKKIL no caminho: %s \n", err)
		return err
	}

	fmt.Printf("Realizado o fechamento do %s \n", strconv.Itoa(pid))

	return nil
}

func OpenMicroSIP(filePath string) error {
	cmd := exec.Command(filePath)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erro ao executar o MicroSIP no caminho: %s \n", err)
		return err
	}

	return nil

}

func GetPIDbyName(processName string) (int, error) {

	cmd := exec.Command("tasklist", "/FO", "CSV", "/NH")
	output, err := cmd.Output()
	if err != nil {
		return 0, nil
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) >= 2 {

			name := strings.Trim(fields[0], "\"")
			pid := strings.Trim(fields[1], "\"")
			if strings.EqualFold(name, processName) {
				return strconv.Atoi(pid)
			}

		}
	}

	return 0, fmt.Errorf("Processo não encontrado: %s", processName)
}
