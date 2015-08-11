package controllers

import (
  "net/http"
  "html/template"
  "github.com/zenazn/goji/web"
  // "../models"
  "../services"
  // "time"
  "fmt"
  "strconv"
)

func BbsView(c web.C, w http.ResponseWriter, r *http.Request) {
  fmt.Println("test")
  // posts := [] models.Post{{1, 1, "Hello GoBBS", time.Now(), time.Now(),time.Time{}}}
  posts := services.GetAllPosts()
  tpl := template.Must(template.ParseFiles("view/bbs.html"))
  tpl.Execute(w, posts)
  fmt.Println(posts)
}

func NewPost(c web.C, w http.ResponseWriter, r *http.Request) {
  tpl := template.Must(template.ParseFiles("view/create.html"))
  tpl.Execute(w, nil)
}

func EditPost(c web.C, w http.ResponseWriter, r *http.Request) {
  id, _:= strconv.ParseInt(c.URLParams["id"], 10, 64)
  post := services.GetPostById(id)
  tpl := template.Must(template.ParseFiles("view/edit.html"))
  tpl.Execute(w, post)
  fmt.Println(post)
}


