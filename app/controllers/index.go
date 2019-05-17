package controller

import (
	"fmt"
	"goblog/app"
)

type IndexController struct {
	app.Controller
}

func (p IndexController) Index() {
	fmt.Println("66666666666666666")
	fmt.Fprint(p.Response, p.Request.RequestURI)
}

func (p IndexController) Info() {
	fmt.Println("777777777777")
	fmt.Fprint(p.Response, p.Request.RequestURI)
}
