package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/agnerft/ListRamais/domain"
)

const (
	baseUrl = "https://basesip.makesystem.com.br"
	path    = "clientes"
	query   = "documento"
)

type ServiceRequest struct {
	httpClient *http.Client
}

func NewServiceCliente() *ServiceRequest {

	return &ServiceRequest{
		httpClient: &http.Client{},
	}
}

func (s *ServiceRequest) RequestJsonCliente(cnpj string) (domain.Cliente, error) {

	response, err := s.httpClient.Get(fmt.Sprintf("%s/%s?%s=%s", baseUrl, path, query, cnpj))

	if err != nil {
		log.Fatal("Erro ao fazer a requisição HTTP:", err)
		return domain.Cliente{}, err
	}
	defer response.Body.Close()

	var cli []domain.Cliente

	// Imprimir o conteúdo do corpo da resposta

	body, err := readBody(response)
	if err != nil {
		return domain.Cliente{}, err
	}

	err = json.Unmarshal([]byte(body), &cli)
	if err != nil {
		log.Fatal("Erro ao decodificar o JSON:", err)
		return domain.Cliente{}, err
	}

	return cli[0], nil
}

func readBody(res *http.Response) (string, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetJson() (*http.Response, []byte, error) {

	resp, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, path))
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	// fmt.Println(string(bodyBytes))

	return resp, bodyBytes, nil
}

func (s *ServiceRequest) PostRamais(url string) (domain.RamalSolo, error) {
	ramalGVC := 7849
	resquestBody := []byte(`{ "cmd" : "sip show peers" }`)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(resquestBody))
	if err != nil {
		return domain.RamalSolo{}, err
	}

	defer resp.Body.Close()

	body := &bytes.Buffer{}

	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return domain.RamalSolo{}, err
	}

	lines := strings.Split(strings.TrimSuffix(body.String(), "\n"), "\n")

	partes := strings.Split(lines[0], ",")

	ramais := make([]int, 0)

	for _, str := range partes {
		parts := strings.Fields(str)

		if len(parts) > 0 {
			primeiraInformacao := parts[0]
			teste := strings.Split(primeiraInformacao, "\\/")
			teste2 := strings.ReplaceAll(teste[0], `"`, "")

			findInt, _ := strconv.Atoi(teste2)

			if findInt != ramalGVC && findInt != 0 && len(strconv.Itoa(findInt)) >= 4 {

				ramais = append(ramais, findInt)
			}

		}

	}

	sliceRamais := intsToRamal(ramais)

	ramal := domain.RamalSolo{
		Ramais: sliceRamais,
	}

	return ramal, nil
}

func intsToRamal(intSlice []int) []domain.Ramal {
	ramalSlice := make([]domain.Ramal, len(intSlice))

	for i, val := range intSlice {
		ramalSlice[i] = domain.Ramal{Sip: val}
	}

	return ramalSlice
}
