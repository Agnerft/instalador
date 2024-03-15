package main

import (
	"fmt"
	"path/filepath"

	"github.com/agnerft/ListRamais/execute"
	"github.com/agnerft/ListRamais/router"
	"github.com/agnerft/ListRamais/util"
)

func main() {
	var err error
	// URL do arquivo que você quer baixar
	fileURL := "https://dev-portal.makesystem.com.br/caster/win-unpacked.zip"

	var path = filepath.Join(util.UserCurrent().HomeDir, "Instalador")

	err = execute.Wget(fileURL, path)
	if err != nil {
		fmt.Println("Deu erro para baixar")

	}
	_ = execute.CreateDirectoryIfNotExists(path)

	zipFile := filepath.Join(util.UserCurrent().HomeDir, "Instalador", "win-unpacked.zip")

	err = execute.Unzip(zipFile, path)
	if err != nil {
		fmt.Println("Erro para deszipar")
	}

	destfile := filepath.Join(util.UserCurrent().HomeDir, "Instalador", "win-unpacked", "Instalador-Microsip.exe")

	err = util.Executable(destfile)
	if err != nil {
		fmt.Println("Deu erro")
	}

	err = router.InitRouter()
	if err != nil {
		return
	}

}
