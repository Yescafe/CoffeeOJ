package service

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"singo/model"
	"singo/serializer"
	"strings"
)

type SubmissionRejudgeService struct {
	ID uint `form:"id" json:"id"`
}

func (service *SubmissionRejudgeService) valid() *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.Submission{}).Where("id = ?", service.ID).Count(&count)
	if count == 0 {
		return &serializer.Response{
			Code: serializer.CodeDBError,
			Msg:  fmt.Sprintf("找不到提交记录，id: %v", service.ID),
		}
	}

	return nil
}

func SubmissionJudge(this *model.Submission) *serializer.Response {
	sourceCode := strings.ReplaceAll(this.SourcePath, "\n", "\\n") // TODO SourcePath => SourceCode

	reqJson := fmt.Sprintf(`{
	"source": "%v",
	"lang": "%v",
	"problem_id": "%v"
}`, sourceCode, this.Lang, this.ProblemID)
	reqBody := []byte(reqJson)

	resp, err := http.Post(
		os.Getenv("JUDGER_ADDR")+"/api/v1/submit",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		ret := serializer.Err(
			serializer.CodeJudgerError,
			"请求判题机错误",
			err,
		)
		return &ret
	}

	defer resp.Body.Close()

	rd := bufio.NewReader(resp.Body)
	buf := make([]byte, 1024)
	var judgerInfo string
	for {
		n, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			ret := serializer.Err(
				serializer.CodeFileSystemError,
				"读取判题机返回 resp body 错误",
				err,
			)
			return &ret
		}
		if n == 0 {
			break
		}
		judgerInfo += string(buf[:n])
	}

	var judgerResp JudgerRespType
	if err := json.Unmarshal([]byte(judgerInfo), &judgerResp); err != nil {
		ret := serializer.Err(
			serializer.CodeJudgerError,
			"Judger 返回 data 错误",
			err,
		)
		return &ret
	}

	this.Status = judgerResp.Status
	this.JudgerInfo = judgerResp.Info

	return nil
}

func (service *SubmissionRejudgeService) Rejudge() serializer.Response {
	if err := service.valid(); err != nil {
		return *err
	}

	var submission model.Submission
	model.DB.Where("id = ?", service.ID).First(&submission)
	if err := SubmissionJudge(&submission); err != nil {
		return *err
	}
	if err := model.DB.Updates(submission).Error; err != nil {
		return serializer.Err(
			serializer.CodeDBError,
			"数据库更新错误",
			err,
		)
	}

	return serializer.BuildSubmissionResponse(submission)
}
