package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// var UsrCurr *user.User

func Executable(filePath string) error {

	cmd := exec.Command(filePath, "/S")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

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

	return 0, err
}

func RemovePath(file string) error {

	err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Verificar se é um arquivo (não um diretório)
		if !info.IsDir() {
			// Excluir o arquivo
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("Erro ao excluir o arquivo %s: %s\n", path, err)
				return err
			}
			fmt.Printf("Arquivo %s excluído com sucesso.\n", path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Erro ao listar arquivos na pasta:", err)
		return err
	}

	return nil
}

func CleanFiles(file string) error {
	dir, err := os.Open(file)
	if err != nil {
		return err
	}

	defer dir.Close()

	files, err := dir.Readdirnames(-1)
	if err != nil {

		return err
	}

	if len(files) == 0 {
		fmt.Println("Pasta vazia")
	} else {

		for _, nameFile := range files {
			err := os.Remove(fmt.Sprintf("%s/%s", file, nameFile))
			if err != nil {
				return err

			}
			fmt.Printf("Removido o arquivo: %s", nameFile)
		}
	}

	return nil
}
