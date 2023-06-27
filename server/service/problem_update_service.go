package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
	"singo/util"
)

type ProblemUpdateService struct {
	ID    uint   `form:"id" json:"id"`
	Title string `form:"title" json:"title" binding:"required"`
	// $MemoLimit MB
	MemoLimit int64 `form:"memo_limit" json:"memo_limit" binding:"required"`
	// $TimeLimit ms
	TimeLimit int64  `form:"time_limit" json:"time_limit" binding:"required"`
	Text      string `form:"text" json:"text" binding:"required"`
}

func (service *ProblemUpdateService) valid() *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.Problem{}).Where("id = ?", service.ID).Count(&count)
	if count == 0 {
		return &serializer.Response{
			Code: serializer.CodeDBError,
			Msg:  fmt.Sprintf("找不到问题，id = %v", service.ID),
		}
	}

	if service.MemoLimit < 16 { // < 16MB
		return &serializer.Response{
			Code: serializer.CodeParamErr,
			Msg:  "内存设置过小，至少为 16 MB",
		}
	}

	if service.TimeLimit < 500 { // < 500ms
		return &serializer.Response{
			Code: serializer.CodeParamErr,
			Msg:  "时间设置过小，至少为 500ms",
		}
	}

	return nil
}

func (service *ProblemUpdateService) Update() serializer.Response {
	if err := service.valid(); err != nil {
		return *err
	}

	var problem model.Problem
	model.DB.Where("id = ?", service.ID).First(&problem)

	problem.Title = service.Title
	problem.MemoLimit = service.MemoLimit
	problem.TimeLimit = service.TimeLimit
	path := problem.Path + "/text.md"
	if err := util.WriteToFile(path, service.Text); err != nil {
		return serializer.Err(
			serializer.CodeFileSystemError,
			fmt.Sprintf("文件 `%v` 写入失败", path),
			err,
		)
	}

	model.DB.Updates(&problem)

	return serializer.BuildProblemResponse(problem)
}
