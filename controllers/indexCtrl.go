package controllers

import (
  "net/http"
  "html/template"
)


type Greet struct {
  Text string
}

func IndexCtrl(w http.ResponseWriter, r *http.Request) {
  greet := Greet{Text: "Hello"}
  tpl := template.Must(template.ParseFiles("view/index.html"))
  tpl.Execute(w, greet)
}


