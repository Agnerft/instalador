package main

import (
	"fmt"

	"github.com/agnerft/ListRamais/services"
)

func main() {

	// var err error
	// fileURL := "https://dev-portal.makesystem.com.br/caster/win-unpacked.zip"

	// err = util.OpenZip(fileURL)
	// if err != nil {
	// 	return
	// }

	// err := router.InitRouter()
	// if err != nil {
	// 	return
	// }

	// var svc services.ServiceRequest

	url := "http://msb2b.gvctelecom.com.br:1127"

	obj, err := services.NewServiceCliente().Encapsule(url)
	if err != nil {
		fmt.Println("Deu bigode")
	}

	for i := range obj {

		fmt.Println(obj[i])
	}

}
