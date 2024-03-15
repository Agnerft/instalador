package execute

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unzip(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return errors.New(err.Error())
	}
	defer reader.Close()

	for _, file := range reader.File {
		targetPath := filepath.Join(destination, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(targetPath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
			return errors.New(err.Error())
		}

		zipFile, err := file.Open()
		if err != nil {
			return errors.New(err.Error())
		}
		defer zipFile.Close()

		targetFile, err := os.Create(targetPath)
		if err != nil {
			return errors.New(err.Error())
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, zipFile); err != nil {
			return errors.New(err.Error())
		}
	}

	return nil
}

func CreateDirectoryIfNotExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Erro ao criar o diretório:", err)
			return errors.New(err.Error())
		}
		fmt.Println("Diretório", dirPath, "criado com sucesso.")
	} else if err != nil {
		fmt.Println("Erro ao verificar o diretório:", err)
		return errors.New(err.Error())
	} else {
		fmt.Println("O diretório", dirPath, "já existe.")
	}
	return nil
}
