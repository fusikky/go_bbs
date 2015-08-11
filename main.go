package main

import (
  _ "github.com/go-sql-driver/mysql"
  "net/http"
  "github.com/zenazn/goji"
  "github.com/zenazn/goji/web"
  "github.com/zenazn/goji/web/middleware"
  // "html/template"
  // "log"
   "./controllers"
   "./services"
  // "fmt"
)

func loginCtrl(w http.ResponseWriter, r *http.Request) {

}

func main() {

  post := web.New()
  goji.Handle("/post/*", post)
  index := web.New()
  goji.Handle("/*", index)

  post.Use(middleware.SubRouter)
  index.Use(middleware.SubRouter)

  index.Get("/", controllers.IndexCtrl)

  post.Get("/bbs", controllers.BbsView)

  post.Post("/create", services.CreatePost)
  post.Get("/edit/:id", controllers.EditPost)
  post.Get("/new", controllers.NewPost)
  post.Post("/update/:id", services.UpdatePost)
  post.Get("/delete/:id", services.DeletePost)

  //http.HandleFunc("/", controllers.IndexCtrl)
  // http.HandleFunc("/bbs", controllers.BbsView)
  // http.HandleFunc("/login", loginCtrl)

  //http.HandleFunc("/post/new", controllers.NewPost)
  //http.HandleFunc("/post/edit/:id", controllers.EditPost)

  // for Post
  //http.HandleFunc("/post/create", services.CreatePost)

  // http.HandleFunc("/post/delete/:id", controllers.DeletePost)

  goji.Serve()
 //  err := http.ListenAndServe(":9090", nil)
 //  fmt.Println(":9090")
 //  if err != nil {
 //    log.Fatal("ListenAndServe:", err)
 //  }
}
