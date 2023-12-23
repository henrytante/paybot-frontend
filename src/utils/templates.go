package utils

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var loadedTemplates *template.Template

func LoadTemplates() error {
	templatesPattern := "views/*.html"
	loadedTemplates = template.New("")

	templateFiles, err := filepath.Glob(templatesPattern)
	if err != nil {
		return err
	}

	loadedTemplates, err = loadedTemplates.ParseFiles(templateFiles...)
	if err != nil {
		return err
	}

	return nil
}

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	err := loadedTemplates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
