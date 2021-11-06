package users

import (
	"github.com/jinzhu/gorm"
	"realworld/common"
)

type UserModel struct {
	ID uint `gorm:"primary_key"`
	Username string `gorm:"column:username"`
	Email string `gorm:"column:email;unique_index"`
	Bio string `gorm:"column:bio;size:1024"`
	Image *string `gorm:"column:image"`
	PasswordHash string `gorm:"column:password;not null"`
}

type FollowModel struct {
	gorm.Model
	Following UserModel
	FollowID uint
	FollowedBy UserModel
	FollowedByID uint
}

// AutoMigrate can migrate schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&FollowModel{})
}
