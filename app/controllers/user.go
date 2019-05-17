package controller

import (
	"fmt"
	"goblog/app"
)

// UserController sss
type UserController struct {
	app.Controller
}

// Do sss
func (p UserController) Login() {
	fmt.Println("I`m UserController")
	fmt.Fprint(p.Response, p.Request.RequestURI)
}
