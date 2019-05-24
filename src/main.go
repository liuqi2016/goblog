package main

import (
	"blog/goblog/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	routers.InitRoute(r)
	r.Run(":80") // listen and serve on 0.0.0.0:80
}
