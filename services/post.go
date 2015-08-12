package services

import (
	"../models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"strconv"
	"time"
)

var engine *xorm.Engine

func init() {
	// db := sql.Open("mysql", "user:password@tpc(host:port)/dbname")
	dbStr := "gobbsuser:gobbs@tcp(localhost:3306)/gobbs"
	var err error
	engine, err = xorm.NewEngine("mysql", dbStr)
	if err != nil {
		log.Fatal("open error: %v", err)
	}
}

func GetPostById(id int64) models.Post {
	post := models.Post{Id: id}
	engine.Get(&post)

	return post
}

func GetAllPosts() []models.Post {
	posts := []models.Post{}
	engine.Find(&posts)

	return posts
}

func CreatePost(c web.C, w http.ResponseWriter, r *http.Request) {
	user := GetUserFromSession(c, w, r)
	post := models.Post{UserId: user.Id, Text: r.FormValue("Text"), CreatedAt: time.Now(), UpdatedAt: time.Now()}

	affected, err := engine.Insert(&post)
	fmt.Println(affected)
	fmt.Println(err)
	fmt.Println(post)

	http.Redirect(w, r, "/post/bbs", 301)

}

func UpdatePost(c web.C, w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(c.URLParams["id"], 10, 64)
	fmt.Println(id)
	post := models.Post{Id: id, Text: r.FormValue("Text"), CreatedAt: time.Now(), UpdatedAt: time.Now()}

	affected, err := engine.Id(id).Update(&post)
	fmt.Println(affected)
	fmt.Println(err)
	fmt.Println(post)

	http.Redirect(w, r, "/post/bbs", 301)

}

func DeletePost(c web.C, w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(c.URLParams["id"], 10, 64)
	post := models.Post{Id: id}
	affected, err := engine.Delete(&post)
	fmt.Println(affected)
	fmt.Println(err)
	fmt.Println(post)

	http.Redirect(w, r, "/post/bbs", 301)
}
