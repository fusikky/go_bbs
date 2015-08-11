package services

import (
    "github.com/zenazn/goji/web"
//    "github.com/wcl48/valval"
//    "github.com/jinzhu/gorm"
//    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "net/http"
//    "encoding/base64"
//     "strings"
//     "html/template"
    "../models"
    "fmt"
    "time"
    "log"
    "strconv"
)

// var tpl *template.Template
// ベーシック認証のIDとパスワード
// const Password = "user:user"

type FormData struct {
  Post models.Post
  Mess string
}

var engine *xorm.Engine

func init(){
  // db := sql.Open("mysql", "user:password@tpc(host:port)/dbname")
  dbStr := "root@tcp(localhost:3306)/gobbs"
  var err error
  engine, err = xorm.NewEngine("mysql", dbStr)
  if err != nil {
    log.Fatal("open error: %v", err)
  }
}

func GetPostById(id int64) models.Post {
  post := models.Post{Id: id}
  engine.Get(&post)
  fmt.Println(post)

  return post
}

func GetAllPosts() [] models.Post {
  posts := [] models.Post{}
  engine.Find(&posts)
  fmt.Println(posts)

  return posts
}

func CreatePost(c web.C, w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "POST":
      fmt.Println("Post")
    case "GET":
      fmt.Println("GET")
  }
  post := models.Post {Text: r.FormValue("Text"),CreatedAt: time.Now(),UpdatedAt: time.Now()}

  affected, err := engine.Insert(&post)
  fmt.Println(affected)
  fmt.Println(err)
  fmt.Println(post)

  http.Redirect(w, r, "/post/bbs", 301)

}

func UpdatePost(c web.C, w http.ResponseWriter, r *http.Request) {

  id, _:= strconv.ParseInt(c.URLParams["id"], 10, 64)
  fmt.Println(id)
  fmt.Println("desu")
  post := models.Post {Id: id, Text: r.FormValue("Text"),CreatedAt: time.Now(),UpdatedAt: time.Now()}

  affected, err := engine.Id(id).Update(&post)
  fmt.Println(affected)
  fmt.Println(err)
  fmt.Println(post)

  http.Redirect(w, r, "/post/bbs", 301)

}

func DeletePost(c web.C, w http.ResponseWriter, r *http.Request) {

  id, _:= strconv.ParseInt(c.URLParams["id"], 10, 64)
  post := models.Post {Id: id}
  affected, err := engine.Insert(&post)
  fmt.Println(affected)
  fmt.Println(err)
  fmt.Println(post)

  http.Redirect(w, r, "/post/bbs", 301)
}

// 
// func UserEdit(c web.C, w http.ResponseWriter, r *http.Request) {
//   User := models.User{}
//   User.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
//   db.Find(&User)
//   tpl = template.Must(template.ParseFiles("view/user/edit.html"))
//   tpl.Execute(w, FormData{User, ""})
// }
// 
// func UserUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
//   User := models.User{}
//   User.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
//   db.Find(&User)
//   User.Name = r.FormValue("Name")
//   if err := models.UserValidate(User); err != nil {
//     var Mess string
//     errs := valval.Errors(err)
//     for _, errInfo := range errs {
//       Mess += fmt.Sprint(errInfo.Error)
//     }
//     tpl = template.Must(template.ParseFiles("view/user/edit.html"))
//     tpl.Execute(w, FormData{User, Mess})
//   } else {
//     db.Save(&User)
//     http.Redirect(w, r, "/user/index", 301)
//   }
// }
// 
// func UserDelete(c web.C, w http.ResponseWriter, r *http.Request) {
//   User := models.User{}
//   User.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
//   db.Delete(&User)
//   http.Redirect(w, r, "/user/index", 301)
// }
// 
// func SuperSecure(c *web.C, h http.Handler) http.Handler {
//   fn := func(w http.ResponseWriter, r *http.Request) {
//     auth := r.Header.Get("Authorization")
//     if !strings.HasPrefix(auth, "Basic ") {
//       pleaseAuth(w)
//       return
//     }
// 
//     password, err := base64.StdEncoding.DecodeString(auth[6:])
//     if err != nil || string(password) != Password {
//       pleaseAuth(w)
//       return
//     }
//     h.ServeHTTP(w, r)
//   }
//   return http.HandlerFunc(fn)
// }
// 
// // Authヘッダを受け付けるための処理
// func pleaseAuth(w http.ResponseWriter) {
//   w.Header().Set("WWW-Authenticate", `Basic realm="Gritter"`)
//   w.WriteHeader(http.StatusUnauthorized)
//   w.Write([]byte("Go away!\n"))
// }


