package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
	"singo/util"
)

type ProblemFetchService struct {
	ID uint `form:"id" json:"id"`
}

type ProblemFetchServiceData struct {
	Title     string `form:"title" json:"title"`
	MemoLimit int64  `form:"memo_limit" json:"memo_limit"`
	TimeLimit int64  `form:"time_limit" json:"time_limit"`
	Text      string `from:"text" json:"text"`
}

func (service *ProblemFetchService) valid() *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.Problem{}).Where("id = ?", service.ID).Count(&count)
	if count == 0 {
		return &serializer.Response{
			Code: serializer.CodeParamErr,
			Msg:  "题目不存在或者已经删除",
		}
	}

	return nil
}

func (service *ProblemFetchService) Fetch() serializer.Response {
	if err := service.valid(); err != nil {
		return *err
	}

	var problem model.Problem
	model.DB.Model(&model.Problem{}).Where("id = ?", service.ID).First(&problem)
	path := problem.Path + "/text.md"
	text, err := util.ReadFromFile(path)
	if err != nil {
		return serializer.Err(
			serializer.CodeFileSystemError,
			fmt.Sprintf("文件 %v 读取失败", path),
			err,
		)
	}

	return serializer.Response{
		Data: ProblemFetchServiceData{
			Title:     problem.Title,
			MemoLimit: problem.MemoLimit,
			TimeLimit: problem.TimeLimit,
			Text:      *text,
		},
	}
}
