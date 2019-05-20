package controller

import (
	"fmt"
	"goblog/app"
	"text/template"
)

type IndexController struct {
	app.Controller
}

func (p IndexController) Index() {
	t, _ := template.ParseFiles("views/index/index.html")
	t.Execute(p.Response, map[string]string{"title": "测试", "time": "1111111"})
}

func (p IndexController) Info() {
	fmt.Println("777777777777")
	fmt.Fprint(p.Response, p.Request.RequestURI)
}
