package v1

import (
	"blog/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result 返回 json格式
func Result(c *gin.Context, code int, m string) {
	if m == "" {
		m = e.GetMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  m,
	})
}
