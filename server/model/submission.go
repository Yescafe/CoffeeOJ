package model

import (
	"gorm.io/gorm"
)

type SubmissionStatus int16

const (
	SubmissionStatusAccepted SubmissionStatus = iota
	SubmissionStatusWrongAnswer
	SubmissionStatusCompilationError
	SubmissionStatusRuntimeError
	SubmissionStatusTimeLimitExceeded
	SubmissionStatusMemoLimitExceeded
	SubmissionStatusUnknownError
)

type Submission struct {
	gorm.Model
	ProblemID  uint
	UserID     uint
	SourcePath string `gorm:"not null;size=1000"`
	Lang       string
	Status     SubmissionStatus
	JudgerInfo string
}

// GetUser 从 ID 获取提交记录
func GetSubmission(ID interface{}) (Submission, error) {
	var submission Submission
	result := DB.First(&submission, ID)
	return submission, result.Error
}
