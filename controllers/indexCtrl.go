package controllers

import (
	"html/template"
	"net/http"
)

func IndexCtrl(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("view/index.html"))
	tpl.Execute(w, nil)
}
