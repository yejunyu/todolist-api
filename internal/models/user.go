package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 表示一个用户账户
type User struct {
	gorm.Model
	// Username 用户名，必须唯一且不能为空
	Username string `gorm:"unique;not null" json:"username" example:"johndoe"`
	// Password 用户密码，在JSON中不显示，存储为哈希值
	Password string `gorm:"not null" json:"-"`
	// Todos 该用户创建的所有待办事项
	Todos []Todo `json:"todos,omitempty"`
}

// BeforeSave 在保存用户之前自动哈希密码
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// 判断是否需要哈希密码：
	// 1. 如果是新记录 (ID为0)
	// 2. 或者，如果是更新记录且 Password 字段被修改了
	if u.ID == 0 || tx.Statement.Changed("Password") {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		// 在这里，明文密码被替换成了哈希值
		u.Password = string(hashedPassword)
	}
	return
}
