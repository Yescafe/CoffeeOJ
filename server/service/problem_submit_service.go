package service

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"singo/serializer"
	"strings"
)

type ProblemSubmitService struct {
	ProblemId  uint   `form:"problem_id" json:"problem_id"`
	SourceCode string `form:"source_code" json:"source_code"`
	Lang       string `form:"lang" json:"lang"`
}

func (service *ProblemSubmitService) Submit() serializer.Response {
	sourceCode := strings.ReplaceAll(service.SourceCode, "\n", "\\n")

	reqJson := fmt.Sprintf(`{
	"source": "%v",
	"lang": "%v",
	"problem_id": "%v"
}`, sourceCode, service.Lang, service.ProblemId)
	fmt.Println(reqJson)
	reqBody := []byte(reqJson)

	resp, err := http.Post(
		os.Getenv("JUDGER_ADDR")+"/api/v1/submit",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return serializer.Err(
			serializer.CodeJudgerError,
			"请求判题机错误",
			err,
		)
	}
	defer resp.Body.Close()

	rd := bufio.NewReader(resp.Body)
	buf := make([]byte, 1024)
	var judgerInfo string
	for {
		n, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			return serializer.Err(
				serializer.CodeFileSystemError,
				"读取判题机返回 resp body 错误",
				err,
			)
		}
		if n == 0 {
			break
		}
		judgerInfo += string(buf[:n])
	}

	return serializer.Response{
		Data: judgerInfo,
	}
}
