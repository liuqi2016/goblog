package v1

import (
	"blog/src/models"
	"blog/src/pkg/e"
	"log"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	// 参数验证
	{
		user := models.User{}
		userName := c.Query("user_name")
		password := c.Query("password")
		email := c.Query("email")
		sex := c.Query("sex")
		birthDay, err := strconv.ParseInt(c.Query("birth_day"), 10, 64)
		var code int
		if err != nil {
			code = e.ERROR_STRTOINT64_FAIL
		} else {
			valid := validation.Validation{}
			valid.Required(userName, "user_name").Message("名称不能为空")
			valid.MaxSize(userName, 100, "user_name").Message("名称最长为60字符")
			valid.Email(email, "email").Message("邮箱格式不正确")
			valid.MinSize(password, 6, "password").Message("密码最小长度为6")
			valid.MaxSize(password, 20, "password").Message("密码最大长度为20")
			code = e.INVALID_PARAMS
			if !valid.HasErrors() {
				if !models.ExistUserByName(userName) {
					code = e.SUCCESS
					user.UserName = userName
					user.Password = password
					user.Sex = sex
					user.Email = email
					user.BirthDay = birthDay
					models.RegisterUser(&user)
				} else {
					code = e.ERROR_EXIST_USER
				}
			} else {
				for _, err := range valid.Errors {
					log.Println(err.Key, err.Message)
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
	}
}

// LoginUser  登录
func LoginUser(c *gin.Context) {

}

// GetUser  获取信息
func GetUser(c *gin.Context) {

}

// LogoutUser  注销
func LogoutUser(c *gin.Context) {

}

// EditUser 编辑
func EditUser(c *gin.Context) {

}
