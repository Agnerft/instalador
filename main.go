package main

import (
	"fmt"

	"github.com/agnerft/ListRamais/services"
)

func main() {

	// var err error
	// fileURL := "https://dev-portal.makesystem.com.br/caster/win-unpacked.zip"

	// err = util.OpenZip(fileURL)
	// if err != nil {
	// 	return
	// }

	// err := router.InitRouter()
	// if err != nil {
	// 	return
	// }

	var svc services.ServiceRequest

	url := "http://umusul.gvctelecom.com.br:1135"

	// obj, err := svc.ObjGVC(url)

	obj, err := svc.Encapsule(url)
	if err != nil {
		fmt.Println("Deu bigode")
	}
	fmt.Println(obj)

	// teste := `"7801\/7801`
	// reStg := regexp.MustCompile(`(.*)\\/(.*)`)
	// reStg := regexp.MustCompile("\\d+")
	// var ramais []string
	// // worldlimpa := strings.Split(words[0], "\\/")
	// // 			worldlimpa = strings.Split(worldlimpa[0], `"`)
	// // 			worldlimpa = strings.Split(worldlimpa[1], " ")

	// for i, s := range obj {
	// 	words := strings.Fields(obj[i])

	// 	if strings.Contains(s, "peers") {
	// 		break
	// 	}

	// 	teste := reStg.FindAllString(words[0], -1)

	// 	// fmt.Println(teste)

	// 	if len(teste) == 2 {
	// 		fmt.Println("primeiro if")

	// 		if teste[0] == teste[1] {
	// 			fmt.Println("segundo if")
	// 			fmt.Println(teste)

	// 			ramais = append(ramais, teste[0])
	// 		}
	// 	} else if len(teste) < 2 {
	// 		fmt.Println("Entrou no else")
	// 		fmt.Println(teste)
	// 		for _, ramal := range teste {
	// 			if ramal == "7849" && len(ramal) == 4 {
	// 				fmt.Println("Ramal de teste GVC")
	// 			}

	// 			ramais = append(ramais, ramal)
	// 		}

	// 	} else {
	// 		fmt.Println(s)
	// 		break
	// 	}
	// }

	// fmt.Println(ramais)

}
