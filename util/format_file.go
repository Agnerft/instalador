package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FormatFile(arquivoEntrada, arquivoSaida string) error {

	// Abre o arquivo de entrada para leitura
	arquivoEntradaObj, err := os.Open(arquivoEntrada)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo de entrada: %s\n Causando o erro %s.", arquivoEntrada, err)
		return err
	}
	defer arquivoEntradaObj.Close()

	// Abre o arquivo de saída para escrita
	arquivoSaidaObj, err := os.Create(arquivoSaida)
	if err != nil {
		fmt.Printf(`Erro ao criar o arquivo de saída: %s \n Causando o erro %s`, arquivoSaida, err)
		return err
	}
	defer arquivoSaidaObj.Close()

	// Cria um scanner para ler o arquivo de entrada linha por linha
	scanner := bufio.NewScanner(arquivoEntradaObj)

	// Indica se a seção Account1 está sendo processada
	estaNaAccount1 := false

	// Loop para ler e processar cada linha
	for scanner.Scan() {
		linha := scanner.Text()

		// Verifica se começou a seção Account1
		if strings.HasPrefix(linha, "[Account1]") {
			estaNaAccount1 = true
		}

		// Se estiver na seção Account1, pula as linhas até encontrar uma linha em branco
		if estaNaAccount1 {
			if linha == "" {
				estaNaAccount1 = false
				continue
			}
			continue
		}

		// Escreve a linha no arquivo de saída
		fmt.Fprintln(arquivoSaidaObj, linha)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo de entrada: %s\n Causando o erro %s", arquivoEntrada, err)
		return err
	}

	fmt.Println("Parte selecionada foi removida e o arquivo foi salvo com sucesso.")

	return nil
}
