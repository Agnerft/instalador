package services

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/agnerft/ListRamais/domain"
)

func (s *ServiceRequest) Encapsule(url string) ([]domain.ObjetoGvc, error) {

	body, err := s.ObjGVC(fmt.Sprintf("%s/%s", url, "asterisk_exec"))
	if err != nil {
		fmt.Println("Deu bigode")
	}
	var worldReal string
	Objs := make([]domain.ObjetoGvc, 0)
	var objGVC domain.ObjetoGvc
	reInt := regexp.MustCompile(`[0-9]`)
	reStg := regexp.MustCompile(`(.*)\\/(.*)`)

	for i, s := range body {

		if strings.Contains(s, "peers") {
			break
		}

		if !strings.Contains(s, "peers") && i != 0 {

			words := strings.Fields(body[i])
			matches := reStg.FindAllStringSubmatch(words[0], -1)
			quant := reInt.FindAllString(words[0], -1)

			for _, match := range matches {

				if match[1] == match[2] && len(quant) == 8 {
					worldReal = match[2]
				}

				worldlimpa := strings.Split(words[0], "\\/")
				worldlimpa = strings.Split(worldlimpa[0], `"`)
				worldlimpa = strings.Split(worldlimpa[1], " ")

				worldReal = worldlimpa[0]
			}

			if words[6] == "UNKNOWN" {
				fmt.Println("banana")
			}

			objGVC = domain.ObjetoGvc{

				NameUsername: worldReal,
				Host:         words[1],
				Dyn:          words[2],
				Forcerport:   words[3],
				Comedia:      words[4],
				ACL:          "null",
				Port:         words[5],
				Status:       fmt.Sprintf("%s%s%s", words[6], words[7], words[8]),
				// Description:  strings.Join(words[8:indiceDescricao], ""),
			}

			Objs = append(Objs, objGVC)
		}

	}

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
