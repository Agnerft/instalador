package main

import (
	"fmt"

	"github.com/agnerft/ListRamais/router"
)

func main() {

	err := router.InitRouter()
	if err != nil {
		fmt.Println("Erro ai inicializar o InitRouter")
	}
}
