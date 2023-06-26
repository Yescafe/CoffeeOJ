package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
)

type SubmissionFetchService struct {
	ID uint `form:"id" json:"id"`
}

func (service *SubmissionFetchService) valid() *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.Submission{}).Where("id = ?", service.ID).Count(&count)
	if count == 0 {
		return &serializer.Response{
			Code: serializer.CodeDBError,
			Msg:  fmt.Sprintf("找不到提交，id: ", service.ID),
		}
	}

	return nil
}

func (service *SubmissionFetchService) Fetch() serializer.Response {
	if err := service.valid(); err != nil {
		return *err
	}

	submission := model.Submission{}
	model.DB.Model(&model.Submission{}).Where("id = ?", service.ID).Model(&submission)

	return serializer.BuildSubmissionResponse(submission)
}
