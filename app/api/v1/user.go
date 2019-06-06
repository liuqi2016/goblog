package v1

import (
	"blog/models"
	"blog/pkg/e"
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// // LoginUser  登录授权
// func LoginUser(c *gin.Context) {

// }

// GetUser  获取信息
func GetUser(c *gin.Context) {

}

// LogoutUser  注销
func LogoutUser(c *gin.Context) {
	//销毁签名

}

// SaveOrEditUser 编辑或者新增
func SaveOrEditUser(c *gin.Context) {
	// 参数验证
	var code = e.SUCCESS
	var msg string
	{
		user := models.User{}
		err := c.BindJSON(&user)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  err.Error(),
			})
			return
		}
		valid := validation.Validation{}
		valid.Required(user.UserName, "user_name").Message("名称不能为空")
		valid.MaxSize(user.UserName, 100, "user_name").Message("名称最长为60字符")
		valid.Email(user.Email, "email").Message("邮箱格式不正确")
		valid.MinSize(user.Password, 6, "password").Message("密码最小长度为6")
		valid.MaxSize(user.Password, 20, "password").Message("密码最大长度为20")
		if !valid.HasErrors() {
			if !models.ExistUserByName(user.UserName, user.ID) {
				if user.ID != 0 {
					err = models.EditUser(&user)
				} else {
					err = models.RegisterUser(&user)
				}
				if err != nil {
					code = 400
					msg = err.Error()
				}
			} else {
				code = e.INVALID_PARAMS
				msg = "该名字已存在"
			}
		} else {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
				msg = err.Message
				break
			}
		}
		if msg == "" {
			msg = e.GetMsg(code)
		}
		Result(c, code, msg)
	}
}
