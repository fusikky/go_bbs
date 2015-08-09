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


