package main

import (
	"fmt"
	"goblog/app"
	controller "goblog/app/controllers"
)

func main() {
	application := app.New()
	application.Get("index", &controller.IndexController{})
	application.Get("user", &controller.UserController{})
	fmt.Printf("%+v", application)
	application.Run(":8080")
}
