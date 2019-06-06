package routers

import (
	"blog/app/api"
	v1 "blog/app/api/v1"
	"blog/middleware/jwt"
	"blog/pkg/setting"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	// r.Use(getMyLog())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth) // 获取token
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	//用户
	{
		// 注册用户
		apiv1.POST("/users", v1.SaveOrEditUser)
		// 更新用户
		apiv1.POST("/users/:id", v1.SaveOrEditUser)
	}
	{
		// //获取标签列表
		// apiv1.GET("/tags", v1.GetTags)
		// //新建标签
		// apiv1.POST("/tags", v1.AddTag)
		// //更新指定标签
		// apiv1.PUT("/tags/:id", v1.EditTag)
		// //删除指定标签
		// apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
