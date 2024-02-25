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

	file, err := os.OpenFile(destFile, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo 2: %v", err)
		return nil, err
	}

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
	// Abrir o arquivo para leitura e escrita
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return err
	}
	defer file.Close()

	file16utf, err := readUTF16LE(file)
	if err != nil {
		fmt.Println("Erro ao ler o file Utf")
	}

	// Substituir a palavra alvo pela palavra substituta
	conteudoModificado := strings.ReplaceAll(string(utf16.Decode(file16utf)), textSearch, newText)

	// Truncar o arquivo para o tamanho do novo conteúdo
	file.Truncate(0)
	file.Seek(0, 0)
	fmt.Println(conteudoModificado)
	writeUTF16LE(file, conteudoModificado)

	fmt.Println("Substituição concluída com sucesso no arquivo:", filepath)
	return nil
}

// readUTF16LE lê o conteúdo de um arquivo UTF-16 LE.
func readUTF16LE(file *os.File) ([]uint16, error) {
	// Lê o BOM (Byte Order Mark) para determinar a ordem de bytes.
	bom := make([]byte, 2)
	_, err := file.Read(bom)
	if err != nil {
		return nil, err
	}

	// Verifica se o arquivo tem a marca de ordem de bytes correta.
	if bom[0] != 0xFF || bom[1] != 0xFE {
		return nil, fmt.Errorf("arquivo não está em formato UTF-16 LE")
	}

	// Lê o conteúdo do arquivo como runes UTF-16.
	var buffer []uint16
	for {
		var u uint16
		err := binary.Read(file, binary.LittleEndian, &u)
		if err != nil {
			break
		}
		buffer = append(buffer, u)
	}

	return buffer, nil
}

// writeUTF16LE escreve o conteúdo no arquivo usando UTF-16 LE.
func writeUTF16LE(file *os.File, content string) {
	// Escreve o BOM (Byte Order Mark) para indicar UTF-16 LE.
	bom := []byte{0xFF, 0xFE}
	file.Write(bom)

	// Converte a string para runes UTF-16.
	runes := utf16.Encode([]rune(content))

	// Escreve os runes no arquivo usando ordem de bytes little-endian.
	for _, u := range runes {
		binary.Write(file, binary.LittleEndian, u)
	}
}
