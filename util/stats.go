package util

import (
	"os"
)

func FileIsExist(caminho string) bool {
	_, err := os.Stat(caminho)

	return !os.IsNotExist(err)
}
