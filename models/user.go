package models

// User 用户表
type User struct {
	Model
	UserName string `gorm:"type:varchar(20);unique_index" json:"user_name"`
	Password string `gorm:"type:varchar(64);not null" json:"password"`
	Email    string `gorm:"type:varchar(20);unique_index" json:"email"`
	Sex      string `gorm:"type:varchar(1)" json:"sex"`
	BirthDay int64  `gorm:"not null" json:"birth_day"`
}

// RegisterUser 注册
func RegisterUser(u *User) (err error) {
	err = db.Create(u).Error
	return
}

// EditUser 编辑
func EditUser(u *User) (err error) {
	err = db.Save(u).Error
	return
}

// CheckAuth 授权
func CheckAuth(username, password string) bool {
	var auth User
	db.Select("id").Where(User{UserName: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}

// ExistUserByName 检验名字是否存在
func ExistUserByName(username string, id uint) (r bool) {
	user := User{}
	if id != 0 {
		db.Select("id").Where("user_name = ? and id <> ?", username, id).First(&user)
	} else {
		db.Select("id").Where("user_name = ?", username).First(&user)
	}
	if user.ID > 0 {
		r = true
	} else {
		r = false
	}
	return
}
