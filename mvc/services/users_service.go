package services

import (
	"github.com/krisDlittleGiant/GoPrj/mvc/domain"
	"github.com/krisDlittleGiant/GoPrj/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
