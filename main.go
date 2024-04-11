package main

import "github.com/agnerft/ListRamais/router"

func main() {

	var err error
	// fileURL := "https://dev-portal.makesystem.com.br/caster/win-unpacked.zip"

	// err = util.OpenZip(fileURL)
	// if err != nil {
	// 	return
	// }

	err = router.InitRouter()
	if err != nil {
		return
	}

}
