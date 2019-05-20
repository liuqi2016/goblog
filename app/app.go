package app

import (
	"fmt"
	"goblog/utils"
	"net/http"
	"reflect"
)

type application struct {
	routes map[string]interface{}
}

func New() *application {
	app := application{
		routes: make(map[string]interface{}),
	}
	return &app
}
func (p *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	controllerName := r.URL.Query().Get("c")
	actionName := r.URL.Query().Get("a")
	actionName = utils.CamelString(actionName)
	if controllerName == "" || actionName == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	route, ok := p.routes[controllerName]
	if !ok {
		http.Error(w, "Controller Not Found", http.StatusNotFound)
		return
	}
	//获取节点的方法集 验证方法是否正确
	v := reflect.ValueOf(route)
	t := v.Type()
	var methods []string //创建一个切片接受所有方法

	for i := 0; i < v.NumMethod(); i++ {
		// fmt.Println(t.Method(i).Name)
		methods = append(methods, t.Method(i).Name)
	}
	var hasMethod bool = false
	for _, method := range methods {

		if method == actionName {
			hasMethod = true
			break
		}
	}
	// fmt.Println(hasMethod)
	if hasMethod == true {
		ele := reflect.ValueOf(route).Elem()
		ele.FieldByName("Request").Set(reflect.ValueOf(r))
		ele.FieldByName("Response").Set(reflect.ValueOf(w))
		ele.MethodByName(actionName).Call([]reflect.Value{})
	} else {
		fmt.Fprintln(w, "404 no page find")
	}
}

func (p *application) printRoutes() {
	for route, controller := range p.routes {
		ele := reflect.ValueOf(controller).Type().String()
		fmt.Printf("%s %s\n", route, ele)
	}
}

func (p *application) Get(route string, controller interface{}) {
	p.routes[route] = controller
}

func (p *application) Run(addr string) error {
	p.printRoutes()
	fmt.Printf("listen on %s\n", addr)
	return http.ListenAndServe(addr, p)
}
