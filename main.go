package main

import "github.com/agnerft/ListRamais/router"

func main() {

	// var err error
	// fileURL := "https://dev-portal.makesystem.com.br/caster/win-unpacked.zip"

	// err = util.OpenZip(fileURL)
	// if err != nil {
	// 	return
	// }

	err := router.InitRouter()
	if err != nil {
		return
	}

	// ramais := domain.RamalSolo{}

	// ramai1 := domain.Ramal{Sip: 7801}
	// ramai2 := domain.Ramal{Sip: 7802}
	// ramai3 := domain.Ramal{Sip: 7803}
	// ramai4 := domain.Ramal{Sip: 7804}

	// ramais.Ramais = append(ramais.Ramais, ramai1, ramai2, ramai3, ramai4)

	// err := util.FileInfos(ramais)
	// if err != nil {

	// 	fmt.Println("Deu bigode")
	// }

}
