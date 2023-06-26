package service

import (
	"singo/model"
	"singo/serializer"
)

type ProblemSubmitService struct {
	ProblemId  uint   `form:"problem_id" json:"problem_id"`
	SourceCode string `form:"source_code" json:"source_code"`
	Lang       string `form:"lang" json:"lang"`
}

type JudgerRespType struct {
	Status model.SubmissionStatus `json:"status"`
	Info   string                 `json:"info"`
}

func (service *ProblemSubmitService) Submit() serializer.Response {
	submission := model.Submission{
		ProblemID:  service.ProblemId,
		UserID:     0,
		SourcePath: service.SourceCode, // TODO save source code to file
		Lang:       service.Lang,
	}
	if err := SubmissionJudge(&submission); err != nil {
		return *err
	}

	if err := model.DB.Create(&submission).Error; err != nil {
		return serializer.Err(
			serializer.CodeDBError,
			"录入提交记录失败",
			err,
		)
	}

	return serializer.BuildSubmissionResponse(submission)
}
