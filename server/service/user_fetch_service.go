package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
)

type UserFetchService struct {
	ID uint `form:"id" json:"id"`
}

func (service *UserFetchService) valid() *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.User{}).Where("id = ?", service.ID).Count(&count)
	if count == 0 {
		return &serializer.Response{
			Code: serializer.CodeDBError,
			Msg:  fmt.Sprintf("找不到用户，id: %v", service.ID),
		}
	}

	return nil
}

func (service *UserFetchService) Fetch() serializer.Response {
	if err := service.valid(); err != nil {
		return *err
	}

	var user model.User
	model.DB.Where("id = ?", service.ID).First(&user)

	return serializer.BuildUserResponse(user)
}
