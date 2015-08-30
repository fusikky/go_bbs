package main

import (
	"./controllers"
	"./services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func main() {

	index := web.New()
	goji.Handle("/*", index)

	index.Use(middleware.SubRouter)

	index.Get("/", controllers.IndexCtrl)

	index.Get("/post/bbs", controllers.BbsView)
	index.Post("/post/create", services.CreatePost)
	index.Get("/post/edit/:id", controllers.EditPost)
	index.Get("/post/new", controllers.NewPost)
	index.Post("/post/update/:id", services.UpdatePost)
	index.Get("/post/delete/:id", services.DeletePost)

	index.Get("/user/new", controllers.NewUser)
	index.Post("/user/create", services.CreateUser)

	index.Get("/user/login", controllers.Login)
	index.Post("/user/login", services.Login)

	goji.Serve()
}
