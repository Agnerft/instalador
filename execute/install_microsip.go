package execute

import (
	"errors"
	"fmt"
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
