package controllers

import (
  "net/http"
  "html/template"
  // "../models"
  "../services"
  // "time"
  "fmt"
)

func BbsCtrl(w http.ResponseWriter, r *http.Request) {
  // posts := [] models.Post{{1, 1, "Hello GoBBS", time.Now(), time.Now(),time.Time{}}}
  posts := services.GetAllPosts()
  tpl := template.Must(template.ParseFiles("view/bbs.html"))
  tpl.Execute(w, posts)
  fmt.Println(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
  tpl := template.Must(template.ParseFiles("view/create.html"))
  tpl.Execute(w, nil)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
  tpl := template.Must(template.ParseFiles("view/edit.html"))
  tpl.Execute(w, nil)
}

func EditPost(w http.ResponseWriter, r *http.Request) {
  fmt.Println("edit post")
  post := services.GetPostById(1)
  tpl := template.Must(template.ParseFiles("view/edit.html"))
  tpl.Execute(w, post)
  fmt.Println(post)
}


