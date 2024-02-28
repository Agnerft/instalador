package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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
	// Fazer uma requisição HTTP para obter os dados JSON
	// var clientes []Cliente
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

	// if len(cli.Clientes) > 0 {
	// 	return domain.Cliente{}, nil
	// }

	return cli[0], nil
}

func (s *ServiceRequest) RequestJsonRamal(url string) (domain.RamaisRegistrados, error) {
	// Fazer uma requisição HTTP para obter os dados JSON

	response, err := s.httpClient.Get(fmt.Sprintf("%s/%s", url, "status_central"))
	if err != nil {
		log.Fatal("Erro ao fazer a requisição HTTP:", err)
		return domain.RamaisRegistrados{}, err
	}

	defer response.Body.Close()

	var ramais domain.RamaisRegistrados

	body, err := readBody(response)
	if err != nil {
		return domain.RamaisRegistrados{}, err
	}

	err = json.Unmarshal([]byte(body), &ramais)
	if err != nil {
		log.Fatal("Erro ao decodificar o JSON:", err)
		return domain.RamaisRegistrados{}, err
	}
	count := 0
	// Imprimir os dados dos ramais registrados
	for _, ramal := range ramais.RamaisRegistrados {
		// fmt.Printf("SIP: %s, IP: %s\n", ramal.Sip, ramal.Ip)
		if ramal.Sip != "" {
			count++
		}

	}

	fmt.Println(count)

	return ramais, nil
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
