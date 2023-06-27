package service

import (
	"fmt"
	"os"
	"singo/model"
	"singo/serializer"
	"singo/util"
)

type ProblemAddService struct {
	CreatorID uint   `form:"creator_id" json:"creator_id"`
	Title     string `form:"title" json:"title" binding:"required"`
	// $MemoLimit MB
	MemoLimit int64 `form:"memo_limit" json:"memo_limit" binding:"required"`
	// $TimeLimit ms
	TimeLimit int64  `form:"time_limit" json:"time_limit" binding:"required"`
	Text      string `form:"text" json:"text" binding:"required"`
}

func (service *ProblemAddService) valid() *serializer.Response {
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

func (service *ProblemAddService) Add() serializer.Response {
	if err := service.valid(); err != nil {
		return *err
	}

	problem := model.Problem{
		CreatorID: service.CreatorID,
		Title:     service.Title,
		MemoLimit: service.MemoLimit,
		TimeLimit: service.TimeLimit,
		Path:      "",
	}

	if err := model.DB.Create(&problem).Error; err != nil {
		return serializer.ParamErr("添加题目失败", err)
	}

	path := fmt.Sprintf("%v/problems/P%v", os.Getenv("DATA_HOME"), problem.ID)
	problem.Path = path
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		model.DB.Delete(&problem)
		return serializer.Err(
			serializer.CodeFileSystemError,
			fmt.Sprintf("%s 创建失败", path),
			err,
		)
	}
	textPath := path + "/text.md"
	if err := util.WriteToFile(textPath, service.Text); err != nil {
		return serializer.Err(
			serializer.CodeFileSystemError,
			fmt.Sprintf("文件 %v 写入失败", textPath),
			err,
		)
	}

	if err := model.DB.Save(&problem).Error; err != nil {
		model.DB.Delete(&problem)
		return serializer.ParamErr("添加题目失败", err)
	}

	return serializer.BuildProblemResponse(problem)
}
