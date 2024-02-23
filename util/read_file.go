package util

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf16"
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

func ReadFile(filePath string) (string, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo 2: %v", err)
		return "", err
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	content := make([]byte, info.Size())

	_, err = file.Read(content)
	if err != nil {
		return "", err
	}

	return string(content), err

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
	file, err := os.OpenFile(destFile, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo 2: %v", err)
		return err
	}

	defer file.Close()

	// Conteúdo a ser adicionado
	novoConteudo := []rune(`[Account1]
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
`)
	utf := utf16.Encode(novoConteudo)
	// binary.Write(file, binary.LittleEndian, uint16(0xFEFF))

	_, numeroLinhas, err := FileForByte(destFile)
	if err != nil {
		log.Fatal("Erro para modificar o arquivo para Byte.", err)
		return err
	}

	if numeroLinhas <= 106 {

		for _, u := range utf {
			binary.Write(file, binary.LittleEndian, u)
		}

	} else {
		fmt.Println("Não precisa add o arquivo")
		return nil
	}

	return nil
}

func FileForByte(destFile string) ([]byte, int, error) {

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

	// Contar as linhas no conteúdo
	numLinhas, _ := contarLinhasNoConteudo(data[:n])

	// Imprimir o número de linhas

	fmt.Printf("Número de linhas no conteúdo do arquivo: %d\n", numLinhas)

	return data, numLinhas, nil
}

func ReplaceLineOfFile(filepath, textSearch, newText string) error {

	file, err := os.OpenFile(filepath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	tempFile, err := os.CreateTemp("", "tempfile")
	if err != nil {
		return err
	}

	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(tempFile)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, textSearch) {
			line = strings.Replace(line, textSearch, newText, -1)
		}
		fmt.Fprintln(writer, line)
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	writer.Flush()

	err = os.Rename(tempFile.Name(), filepath)
	if err != nil {
		return err
	}
	return nil
}
