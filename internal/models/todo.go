package models

import "gorm.io/gorm"

// Todo 表示一个待办事项
type Todo struct {
	gorm.Model
	// Title 待办事项的标题
	Title string `gorm:"not null" json:"title" example:"完成项目文档"`
	// Status 待办事项的完成状态，false表示未完成，true表示已完成
	Status bool `gorm:"default:false" json:"status" example:"false"`
	// UserId 创建该待办事项的用户ID
	UserId uint `gorm:"not null" json:"uid" example:"1"`
}
