package main

import (
	"fmt"
	"path/filepath"

	"github.com/agnerft/ListRamais/util"
)

func main() {
	destFileConfigMicrosip := filepath.Join(util.UserCurrent().HomeDir, "AppData", "Roaming", "MicroSIP", "microsip.ini")
	err := util.ReplaceLineOfFile(destFileConfigMicrosip, "accountId=0", "accountId=1")
	if err != nil {
		fmt.Println("Deu erro")
	}

}
