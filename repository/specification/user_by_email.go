package specification

import (
	"github.com/dailoi280702/vrs-general-service/type/model"
	"gorm.io/gorm"
)

type userByEmail struct {
	email string
}

func UserByEmail(email string) I {
	return &userByEmail{
		email: email,
	}
}

func (u *userByEmail) GormQuery(db *gorm.DB) *gorm.DB {
	return db.Model(&model.User{}).Where("email = ?", u.email)
}
