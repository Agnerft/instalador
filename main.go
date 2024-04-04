package main

import (
	"fmt"

	"github.com/agnerft/ListRamais/services"
)

func main() {
	var err error
	// // URL do arquivo que você quer baixar
	// fileURL := "https://dev-portal.makesystem.com.br/caster/win-unpacked.zip"

	// var path = filepath.Join(util.UserCurrent().HomeDir, "Instalador")

	// err = execute.Wget(fileURL, path)
	// if err != nil {
	// 	fmt.Println("Deu erro para baixar")

	// }
	// _ = execute.CreateDirectoryIfNotExists(path)

	// zipFile := filepath.Join(util.UserCurrent().HomeDir, "Instalador", "win-unpacked.zip")

	// err = execute.Unzip(zipFile, path)
	// if err != nil {
	// 	fmt.Println("Erro para deszipar")
	// }

	// destfile := filepath.Join(util.UserCurrent().HomeDir, "Instalador", "win-unpacked", "Instalador-Microsip.exe")

	// err = util.Executable(destfile)
	// if err != nil {
	// 	fmt.Println("Deu erro")
	// }

	// err = router.InitRouter()
	// if err != nil {
	// 	return
	// }
	// var extract = map[string]string{}
	url := "http://mscelular.gvctelecom.com.br:1133/asterisk_exec"

	res, err := services.NewServiceCliente().PostRamais(url)
	if err != nil {
		fmt.Println("Deu ruim")
	}

	// partes := strings.Split(res[0], ",")

	// ramais := make([]string, 0)

	// for _, str := range partes {
	// 	parts := strings.Fields(str)

	// 	// Selecionar apenas o primeiro elemento
	// 	if len(parts) > 0 {
	// 		primeiraInformacao := parts[0]
	// 		ramais = append(ramais, primeiraInformacao)
	// 	}

	// }

	// fmt.Println(partes[1])

	// for i, s := range res {
	// 	res[i] = strings.ReplaceAll(s, `"`, ";")
	// }

	fmt.Println(res)

}
