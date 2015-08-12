package services

import (
	"../models"
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
)

func Login(c web.C, w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "sessionUser")
	session.Values["Name"] = r.FormValue("Name")
	user := GetUserByName(session.Values["Name"].(string))
	if user.Id != 0 {
		session.Values["Id"] = user.Id
		session.Save(r, w)
	} else {
		session.Values["Name"] = "げすと"
		session.Save(r, w)
	}

	http.Redirect(w, r, "/post/bbs", 301)
}

func GetUserFromSession(c web.C, w http.ResponseWriter, r *http.Request) models.User {
	session, _ := store.Get(r, "sessionUser")

	user := models.User{Id: session.Values["Id"].(int64), Name: session.Values["Name"].(string)}
	fmt.Println(session.Values["Id"])
	fmt.Println(session.Values["Name"])

	return user
}
