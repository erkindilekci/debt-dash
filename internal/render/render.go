package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templateName string, data any) {
	var tmpl *template.Template
	var err error

	_, inMap := templateCache[templateName]
	if !inMap {
		err = createTemplateCache(templateName)
		if err != nil {
			Catch(err)
		}
	}

	tmpl = templateCache[templateName]
	err = tmpl.Execute(w, data)
	if err != nil {
		Catch(err)
	}
}

func createTemplateCache(templateName string) error {
	templates := []string{
		"./templates/base.layout.gohtml",
		fmt.Sprintf("./templates/%s", templateName),
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	templateCache[templateName] = tmpl

	return nil
}

func Catch(err error) {
	if err != nil {
		log.Println(err)
	}
}
