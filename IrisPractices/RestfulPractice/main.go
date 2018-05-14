package main

import (
	"github.com/kataras/iris"
	"github.com/Humbertzhang/goweb_practice/IrisPractices/RestfulPractice/routes"
)

func main() {
	app := iris.New()

	app.Post("/register", routes.Register)
	app.Post("/login", routes.Login)
	app.Post("/article", routes.PostArticle)
	app.Get("/article/{postid:int}", routes.GetOneArticle)
	app.Get("/articles/{uid:long}", routes.GetMyArticles)

	app.Run(iris.Addr(":8080"))
}
