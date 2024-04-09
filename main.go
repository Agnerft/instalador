package main

import (
	"github.com/agnerft/ListRamais/router"
	"github.com/agnerft/ListRamais/util"
)

func main() {

	var err error
	fileURL := "https://dev-portal.makesystem.com.br/caster/win-unpacked.zip"

	err = util.OpenZip(fileURL)
	if err != nil {
		return
	}

	err = router.InitRouter()
	if err != nil {
		return
	}

	// confiMicroSIP := filepath.Join(util.UserCurrent().HomeDir,
	// 	"AppData",
	// 	"Roaming",
	// 	"MicroSIP")

	// dir, err := os.Open(confiMicroSIP)
	// if err != nil {
	// 	fmt.Println("Deu erro para ler")
	// }

	// defer dir.Close()

	// files, err := dir.Readdirnames(-1)
	// if err != nil {

	// 	fmt.Println("Erro ao ler os nomes dos arquivos.")
	// }

	// if len(files) == 0 {
	// 	fmt.Println("Pasta vazia")
	// } else {
	// 	fmt.Println("Arquivos:")
	// 	for _, nameFile := range files {
	// 		err := os.Remove(fmt.Sprintf("%s/%s", confiMicroSIP, nameFile))
	// 		if err != nil {
	// 			fmt.Println("Deu erro pra excluir")
	// 		}
	// 		fmt.Printf("Removido o arquivo: %s", nameFile)
	// 	}
	// }

}
