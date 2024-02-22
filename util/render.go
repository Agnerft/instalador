package util

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, any any, templateFile string) {

	template, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, "Erro ao renderizar o HTML 1", http.StatusInternalServerError)
		log.Println("Erro:", err)
		return
	}

	err = template.Execute(w, any)
	if err != nil {
		http.Error(w, "Erro ao renderizar o HTML 2", http.StatusInternalServerError)
		log.Println("Erro:", err)
		return
	}
}
