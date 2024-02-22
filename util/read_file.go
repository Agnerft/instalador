package util

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

// func ReadFile(filePath, novoValor string, numeroLinha int) error {

// 	// Ler o arquivo INI existente
// 	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
// 	if err != nil {
// 		log.Fatalf("Erro ao abrir o arquivo: %v", err)
// 	}
// 	defer file.Close()

// 	// fmt.Println(novoValor)
// 	conteudo, err := io.ReadAll(file)
// 	if err != nil {
// 		log.Fatalf("Erro ao ler o conteúdo do arquivo 1: %v", err)
// 	}

// 	// Converte o conteúdo para uma string
// 	conteudoArquivo := string(conteudo)

// 	linhas := strings.Split(conteudoArquivo, "\n")

// 	if numeroLinha > 0 && numeroLinha < len(linhas) {
// 		linhas[numeroLinha-1] = novoValor
// 	}

// 	novoConteudoArquivo := strings.Join(linhas, "\n")

// 	_, err = file.Seek(0, 0)
// 	if err != nil {
// 		log.Fatalf("Erro ao mover o ponteiro: %v", err)

// 	}

// 	_, err = file.WriteString(novoConteudoArquivo)
// 	if err != nil {
// 		log.Fatalf("Erro ao salvar novo conteudo: %v", err)
// 	}

// 	err = file.Truncate(int64(len(novoConteudoArquivo)))
// 	if err != nil {
// 		log.Fatalf("Erro ao truncar: %v", err)
// 	}
// 	// // Criar um scanner para ler o conteúdo do arquivo linha por linha
// 	// fmt.Println(novoConteudoArquivo)

// 	return nil
// }

func ReadFile(filePath, oldText, newText string, position int) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)

	}

	// SEGUE

	teste := strings.Split(string(data), "\n")

	if string(teste[position]) == oldText {

		txt := strings.Replace(string(data), oldText, newText, 1)
		err = os.WriteFile(filePath, []byte(txt), 0755)
		if err != nil {
			panic(err)

		}
	} else {
		fmt.Printf("O texto que está no %s é igual ao %s, por isso não precisa gravar.\n", teste[position], newText)
	}

	return nil

}

func contarLinhasNoConteudo(data []byte) (int, error) {
	// Inicializar o contador de linhas
	numLinhas := 0

	// Criar um scanner para contar as linhas no conteúdo
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		numLinhas++
	}

	// Verificar por erros durante o scanner
	if err := scanner.Err(); err != nil {
		log.Fatal("Erro ao contar as linhas no conteúdo:", err)
	}

	return numLinhas, nil
}

func AdicionarConfiguracao(destFile string) error {
	// Abrir o arquivo em modo de escrita, criando-o se não existir
	// file, err := os.OpenFile(destFile, os.O_RDWR|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatalf("Erro ao ler o conteúdo do arquivo 2: %v", err)
	// 	return err
	// }

	// defer file.Close()

	data, err := os.ReadFile(destFile)
	if err != nil {
		panic(err)

	}

	// Conteúdo a ser adicionado
	novoConteudo := `[Account1]
label=
server=
proxy=
domain=
username=
password=
authID=
displayName=
dialingPrefix=
dialPlan=
hideCID=0
voicemailNumber=
transport=
publicAddr=
SRTP=
registerRefresh=300
keepAlive=15
publish=0
ICE=0
allowRewrite=0
disableSessionTimer=0
`

	_, numeroLinhas, err := fileForByte(destFile)
	if err != nil {
		log.Fatal("Erro para modificar o arquivo para Byte.", err)
		return err
	}

	if numeroLinhas <= 106 {
		// Add novo conteudo
		data = append(data, []byte(novoConteudo)...)
		// Escrever o novo conteúdo no arquivo
		err = os.WriteFile(destFile, data, 0755)
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("Não precisa add o arquivo")
	}

	return nil
}

func fileForByte(destFile string) ([]byte, int, error) {

	// Abrir o arquivo em modo de escrita, criando-o se não existir
	file, err := os.OpenFile(destFile, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo 2: %v", err)
		return nil, 0, err
	}
	// Estatísticas sobre o arquivo
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("Erro ao obter informações sobre o arquivo:", err)
	}

	// Leia o conteúdo do arquivo
	data := make([]byte, fileInfo.Size())
	n, err := file.Read(data)
	if err != nil {
		log.Fatal("Erro ao ler o arquivo:", err)
		return nil, 0, err
	}

	// Imprima o conteúdo lido
	// fmt.Printf("Conteúdo do arquivo:\n%s\n", data[:n])

	// Contar as linhas no conteúdo
	numLinhas, _ := contarLinhasNoConteudo(data[:n])

	// Imprimir o número de linhas

	fmt.Printf("Número de linhas no conteúdo do arquivo: %d\n", numLinhas)

	return data, numLinhas, nil
}
