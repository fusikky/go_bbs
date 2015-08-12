package controllers

import (
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
	// "../models"
	"../services"
	// "time"
	"fmt"
	"strconv"
)

func NewUser(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Println("NewUser")
	tpl := template.Must(template.ParseFiles("view/user/create.html"))
	tpl.Execute(w, nil)
}

func EditUser(c web.C, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(c.URLParams["id"], 10, 64)
	post := services.GetPostById(id)
	tpl := template.Must(template.ParseFiles("view/user/edit.html"))
	tpl.Execute(w, post)
	//	fmt.Println(post)
}
