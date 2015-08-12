package controllers

import (
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
)

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("view/user/login.html"))
	tpl.Execute(w, nil)
}
