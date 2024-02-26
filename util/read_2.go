package util

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"unicode/utf16"

	"github.com/agnerft/ListRamais/domain"
	"github.com/saintfish/chardet"
)

type IniFile struct {
	Sections      map[string]map[string]string
	PathFile      string
	regexpSection *regexp.Regexp
	regexpValues  *regexp.Regexp
	lineOrder     []string
}

func NewIniFile(filename string) *IniFile {
	return &IniFile{
		PathFile:      filename,
		Sections:      make(map[string]map[string]string),
		regexpSection: regexp.MustCompile(`\[(.*?)\]`),
		regexpValues:  regexp.MustCompile(`(.*?)=(.*)`),
		lineOrder:     make([]string, 0),
	}
}

func (ini *IniFile) IsSection(line string) bool {
	return ini.regexpSection.MatchString(line)
}

func (ini *IniFile) ExistsSection(section string) bool {
	_, ok := ini.Sections[section]
	return ok
}
func (ini *IniFile) ExistsKey(section, key string) bool {
	_, ok := ini.Sections[section][key]
	return ok
}

func (ini *IniFile) GetSectionText(line string) string {
	values := ini.regexpSection.FindStringSubmatch(line)
	if len(values) > 0 {
		return values[1]
	} else {
		return ""
	}
}

func (ini *IniFile) Readini() error {
	ini.converterFileUTF16ToUTF8()
	file, err := os.Open(ini.PathFile)
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(file)
	var section string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if ini.IsSection(line) {
			section = ini.GetSectionText(line)
			ini.lineOrder = append(ini.lineOrder, section)
			ini.Sections[section] = make(map[string]string)
		} else {
			values := ini.regexpValues.FindStringSubmatch(line)
			if len(values) > 0 {
				ini.lineOrder = append(ini.lineOrder, values[1])
				ini.Sections[section][values[1]] = values[2]
			}
		}
	}
	return nil
}

func (ini *IniFile) UpdateSection(section string, key string, value string) {
	ini.Sections[section][key] = value
}

func (ini *IniFile) UpdateBatchSection(section string, values map[string]string) {
	for key, value := range values {
		ini.Sections[section][key] = value
	}
}

func (ini *IniFile) AddSection(section string) {
	if ini.ExistsSection(section) {
		return
	}
	ini.lineOrder = append(ini.lineOrder, section)
	ini.Sections[section] = make(map[string]string)
}

func (ini *IniFile) AddValueToSection(section string, key string, value string) {

	if ini.ExistsKey(section, key) {
		return
	}

	ini.lineOrder = append(ini.lineOrder, key)
	ini.Sections[section][key] = value
}

func (ini *IniFile) WriteIni() error {
	ini.converterFileUTF16ToUTF8()
	file, err := os.OpenFile(ini.PathFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	var asection string
	for _, section := range ini.lineOrder {
		if ini.ExistsSection(section) {
			asection = section
			file.WriteString("[" + section + "]\n")
		} else {
			value := ini.Sections[asection][section]

			file.WriteString(fmt.Sprintf("%s=%s\n", section, value))
		}
	}
	return nil
}

func (ini *IniFile) converterFileUTF16ToUTF8() {
	data, err := os.ReadFile(ini.PathFile)
	if err != nil {
		return
	}
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(data)
	if err != nil {
		return
	}
	if result.Charset == "UTF-16LE" {
		utf16data := make([]uint16, len(data)/2)
		for i := 0; i < len(data); i += 2 {
			utf16data[i/2] = binary.LittleEndian.Uint16(data[i:])
		}
		utf8data := string(utf16.Decode(utf16data))
		os.WriteFile(ini.PathFile, []byte(utf8data), 0644)
	}

}

func (ini *IniFile) AddSectionAccount(section string, Cfg domain.Config) {
	// if ini.ExistsSection(section) {
	var numero int
	// Cfg := *domain.NewConfig()
	regex := regexp.MustCompile(`(\d+)`)
	resInt := regex.FindStringSubmatch(section)
	numero, _ = strconv.Atoi(resInt[0])

	assection := fmt.Sprintf("%s%d", "Account", numero+1)
	ini.AddSection(assection)

	structValue := reflect.ValueOf(Cfg)
	structType := reflect.TypeOf(Cfg)

	for i := 0; i < structValue.NumField(); i++ {
		header := structType.Field(i).Name
		value := structValue.Field(i).Interface()
		fmt.Println(header)
		fmt.Println(value)

		ini.AddValueToSection(assection, header, fmt.Sprintf("%v", value))
	}

	return
	// }
}
