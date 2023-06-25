package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
)

type ProblemDeleteService struct {
	ID uint `form:"id" json:"id"`
}

func (service *ProblemDeleteService) valid() *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.Problem{}).Where("id = ?", service.ID).Count(&count)
	if count == 0 {
		return &serializer.Response{
			Code: serializer.CodeParamErr,
			Msg:  fmt.Sprintf("找不到题目，id = %v", service.ID),
		}
	}

	return nil
}

func (service *ProblemDeleteService) Delete() serializer.Response {
	if err := service.valid(); err != nil {
		return *err
	}

	var problem model.Problem
	model.DB.Model(&model.Problem{}).Where("id = ?", service.ID).First(&problem)

	if err := model.DB.Delete(&problem).Error; err != nil {
		return serializer.Err(
			serializer.CodeDBError,
			"删除失败",
			err,
		)
	}

	return serializer.Response{
		Msg: "删除成功",
	}
}
