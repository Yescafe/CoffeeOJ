package service

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"singo/model"
	"singo/serializer"
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
	fi, err := os.Open(path)
	if err != nil {
		return serializer.Err(
			serializer.CodeFileSystemError,
			fmt.Sprintf("无法打开 %s", path),
			err,
		)
	}
	defer fi.Close()

	rd := bufio.NewReader(fi)
	buf := make([]byte, 1024)
	var text string

	for {
		n, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			return serializer.Err(
				serializer.CodeFileSystemError,
				"读取失败",
				err,
			)
		}
		if n == 0 {
			break
		}
		text += string(buf[:n])
	}

	return serializer.Response{
		Data: ProblemFetchServiceData{
			Title:     problem.Title,
			MemoLimit: problem.MemoLimit,
			TimeLimit: problem.TimeLimit,
			Text:      text,
		},
	}
}
