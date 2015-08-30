package services

import (
	"../models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/sessions"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"strconv"
	"time"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func init() {
	dbStr := "gobbsuser:gobbs@tcp(localhost:3306)/gobbs"
	var err error
	engine, err = xorm.NewEngine("mysql", dbStr)
	if err != nil {
		log.Fatal("open error: %v", err)
	}
}

func GetUserById(id int64) models.User {
	user := models.User{Id: id}
	engine.Get(&user)
	fmt.Println(user)

	return user
}

func GetUserByName(name string) models.User {
	user := models.User{Name: name}
	affected, err := engine.Get(&user)
	fmt.Println(affected)
	fmt.Println(err)
	fmt.Println(user)

	return user
}

func GetAllusers() []models.User {
	users := []models.User{}
	engine.Find(&users)

	return users
}

func CreateUser(c web.C, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Println("Post")
	case "GET":
		fmt.Println("GET")
	}
	user := models.User{Name: r.FormValue("Name"), Password: r.FormValue("Password"), CreatedAt: time.Now(), UpdatedAt: time.Now()}

	affected, err := engine.Insert(&user)
	fmt.Println(affected)
	fmt.Println(err)
	fmt.Println(user)

	session, _ := store.Get(r, "sessionUser")
	session.Values["Name"] = user.Name
	session.Values["Id"] = user.Id
	session.Save(r, w)

	http.Redirect(w, r, "/post/bbs", 301)

}

func UpdateUser(c web.C, w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(c.URLParams["id"], 10, 64)
	user := models.User{Id: id, Name: r.FormValue("Name"), UpdatedAt: time.Now()}

	affected, err := engine.Id(id).Update(&user)
	fmt.Println(affected)
	fmt.Println(err)
	fmt.Println(user)

	http.Redirect(w, r, "/post/bbs", 301)

}

func DeleteUser(c web.C, w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(c.URLParams["id"], 10, 64)
	user := models.User{Id: id}
	affected, err := engine.Delete(&user)
	fmt.Println(affected)
	fmt.Println(err)
	fmt.Println(user)

	http.Redirect(w, r, "/post/bbs", 301)
}
