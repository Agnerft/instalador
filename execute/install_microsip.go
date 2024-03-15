package execute

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

func DownloadGeneric(url string, destination string) error {

	requisicao, err := grab.NewRequest(destination, url)
	if err != nil {
		return err
	}

	downCli := grab.NewClient()
	response := downCli.Do(requisicao)

	monitorandoDown := time.NewTicker(1000 * time.Millisecond)
	defer monitorandoDown.Stop()

Loop:

	for {
		select {
		case <-monitorandoDown.C:
			fmt.Printf("\rProgresso: %.2f%% concluído", response.Progress()*100)

		case <-response.Done:
			break Loop
		}
	}

	if err := response.Err(); err != nil {
		return errors.New(err.Error())
	}

	fmt.Printf("\rDownload concluído: %s -> %s\n", url, response.Filename)
	return nil
}

func Wget(fileURL string, downloadDir string) error {
	// URL do arquivo que você quer baixar

	// Crie um cliente HTTP
	client := &http.Client{}

	// Faça uma solicitação GET para a URL do arquivo
	resp, err := client.Get(fileURL)
	if err != nil {
		fmt.Println("Erro ao fazer solicitação HTTP:", err)
		return err
	}
	defer resp.Body.Close()

	// Certifique-se de que o diretório de download existe
	err = os.MkdirAll(downloadDir, os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar diretório:", err)
		return err
	}

	// Determine o nome do arquivo a partir do URL
	fileName := filepath.Base(fileURL)

	// Crie o caminho completo do arquivo local
	localFilePath := filepath.Join(downloadDir, fileName)

	// Crie um arquivo local onde você deseja salvar o conteúdo
	out, err := os.Create(localFilePath)
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return err
	}
	defer out.Close()

	// Copie o conteúdo do corpo da resposta para o arquivo local
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Erro ao copiar conteúdo para o arquivo:", err)
		return err
	}

	fmt.Println("Arquivo baixado com sucesso em:", localFilePath)

	return nil
}
