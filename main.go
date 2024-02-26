package main

import "github.com/agnerft/ListRamais/router"

func main() {

	// var destFileConfigMicrosip = "MicroSIP.ini"

	// ini := util.NewIniFile(destFileConfigMicrosip)

	// err := ini.Readini()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// ini.AddSectionAccount("Account7")

	// err = ini.WriteIni()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("Cabo")
	err := router.InitRouter()
	if err != nil {
		return
	}

}
