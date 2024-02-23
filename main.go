package main

import (
	"fmt"
	"path/filepath"

	"github.com/agnerft/ListRamais/util"
)

var ()

func main() {
	var destFileConfigMicrosip = filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "microsip.ini")
	// router.InitRouter()

	// linhas, err := util.ReadFile(destFileConfigMicrosip)
	// if err != nil {
	// 	fmt.Printf("erro nas linhas")
	// }

	// fmt.Println(linhas)

	err := util.ReplaceLineOfFile(destFileConfigMicrosip, "accountId=0", "accountId=1")
	if err != nil {
		fmt.Println("Deu erro")
	}

}
