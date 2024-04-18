package services

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/agnerft/ListRamais/domain"
)

func (s *ServiceRequest) Encapsule(url string) ([]domain.ObjetoGvc, error) {

	// url := "http://umusul.gvctelecom.com.br:1135"

	bod, err := s.ObjGVC(fmt.Sprintf("%s/%s", url, "asterisk_exec"))
	if err != nil {
		fmt.Println("Deu bigode")
	}

	Objs := make([]domain.ObjetoGvc, 0)

	camposCabecalho := strings.Fields(bod[0])
	// indiceDescricao := len(camposCabecalho) - 1

	for i, s := range bod {

		if !strings.Contains(s, "peers") && i != 0 {
			words := strings.Fields(bod[i])

			if len(words) != len(camposCabecalho) {
				continue
			}

			indexBarra := strings.Index(words[0], "\\/")
			if indexBarra != -1 {
				words[0] = words[0][:indexBarra]
			}

			// fmt.Printf("%s -> linha %d \n", words, i)
			obj := domain.ObjetoGvc{
				// NameUsername: strings.ReplaceAll(words[0], `"`, ""),
				NameUsername: words[0],
				Host:         words[1],
				Dyn:          words[2],
				Forcerport:   words[3],
				Comedia:      words[4],
				ACL:          "null",
				Port:         words[5],
				Status:       fmt.Sprintf("%s%s%s", words[6], words[7], words[8]),
				// Description:  strings.Join(words[8:indiceDescricao], ""),
			}

			Objs = append(Objs, obj)
			// fmt.Println(Objs)
		}

	}

	// fmt.Println(Objs)

	return Objs, nil
}

func (s *ServiceRequest) ObjGVC(url string) ([]string, error) {
	resquestBody := []byte(`{ "cmd" : "sip show peers" }`)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(resquestBody))
	if err != nil {
		return []string{}, err
	}

	defer resp.Body.Close()

	body := &bytes.Buffer{}

	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(strings.TrimSuffix(body.String(), "\n"), "\n")
	partes := strings.Split(lines[0], ",")

	return partes, nil
}
