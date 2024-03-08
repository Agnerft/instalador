package main

import (
	"fmt"

	"github.com/agnerft/ListRamais/router"
	"github.com/agnerft/ListRamais/util"
)

func main() {
	var err error
	var fileExecutavel = "win-unpacked/Instalador-Microsip.exe"

	err = util.Executable(fileExecutavel)
	if err != nil {
		fmt.Println("Deu erro")
	}

	err = router.InitRouter()
	if err != nil {
		return
		fmt.Println("Deu erro aqui")
	}

}
