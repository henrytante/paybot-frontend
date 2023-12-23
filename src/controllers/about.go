package controllers

import (
	"net/http"
	"webapp/src/utils"
)


func About(w http.ResponseWriter, r *http.Request)  {
	utils.RenderTemplate(w, "about.html", nil)
}