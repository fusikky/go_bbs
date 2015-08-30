package controllers

import (
	"../models"
	"../services"
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
	"strconv"
)

type Bbs struct {
	User  models.User
	Posts []models.Post
}

func BbsView(c web.C, w http.ResponseWriter, r *http.Request) {
	posts := services.GetAllPosts()
	tpl := template.Must(template.ParseFiles("view/bbs.html"))
	user := services.GetUserFromSession(c, w, r)
	tpl.Execute(w, Bbs{User: user, Posts: posts})
}

func NewPost(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("view/create.html"))
	tpl.Execute(w, nil)
}

func EditPost(c web.C, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(c.URLParams["id"], 10, 64)
	post := services.GetPostById(id)
	tpl := template.Must(template.ParseFiles("view/edit.html"))
	tpl.Execute(w, post)
}
