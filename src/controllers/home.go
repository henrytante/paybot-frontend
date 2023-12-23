package controllers

import (
	"html/template"
	"log"
	"net/http"
	
)

func Home(w http.ResponseWriter, r *http.Request)  {
	tmpl, err := template.ParseFiles("views/home.html", "views/navbar.html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil{
		log.Fatal(err)
	}
}