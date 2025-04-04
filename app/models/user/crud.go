package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

func (u *User) Create() (err error) {
	if err = model.DB.Create(&u).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
