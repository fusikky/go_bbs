package main

import (
  _ "github.com/go-sql-driver/mysql"
  "net/http"
  // "html/template"
  "log"
  "./controllers"
  "./services"
  "fmt"
)

func loginCtrl(w http.ResponseWriter, r *http.Request) {

}

func main() {
  http.HandleFunc("/", controllers.IndexCtrl)
  http.HandleFunc("/bbs", controllers.BbsCtrl)
  http.HandleFunc("/login", loginCtrl)

  http.HandleFunc("/post/new", services.CreatePost)

  err := http.ListenAndServe(":9090", nil)
  fmt.Println(":9090")
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
