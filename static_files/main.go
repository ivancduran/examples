package main

import (
	"github.com/iris-contrib/template/html"
	"github.com/kataras/iris"
)

type page struct {
	Title string
}

func main() {
	iris.UseTemplate(html.New()).Directory("./templates/web/default", ".html")
	iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	})
	iris.Static("/css", "./resources/css", 1)
	iris.Static("/js", "./resources/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("something.html", page{Title: "Home"})
	})

	iris.Listen(":8080")
}
