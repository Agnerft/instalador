package util

import (
	"fmt"
	"path/filepath"

	"github.com/agnerft/ListRamais/execute"
)

func OpenZip(fileURL string) error {

	var err error

	var path = filepath.Join(UserCurrent().HomeDir,
		"Instalador")

	err = execute.Wget(fileURL, path)
	if err != nil {
		return fmt.Errorf("%s -> %s", "Deu erro para baixar", err)

	}
	_ = execute.CreateDirectoryIfNotExists(path)

	zipFile := filepath.Join(UserCurrent().HomeDir,
		"Instalador",
		"win-unpacked.zip")

	err = execute.Unzip(zipFile, path)
	if err != nil {
		return fmt.Errorf("%s -> %s", "Erro para deszipar", err)
	}

	destfile := filepath.Join(UserCurrent().HomeDir,
		"Instalador",
		"win-unpacked",
		"Instalador-Microsip.exe")

	err = Executable(destfile)
	if err != nil {
		return fmt.Errorf("%s -> %s", "Erro ao executar", err)
	}
	return nil
}
