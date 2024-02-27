package main

import "github.com/agnerft/ListRamais/router"

func main() {

	err := router.InitRouter()
	if err != nil {
		return
	}

}
