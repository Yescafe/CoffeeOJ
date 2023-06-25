package model

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	ProblemID  uint
	UserID     uint
	SourcePath string `gorm:"not null;size:1000"`
}

// GetUser 从 ID 获取提交记录
func GetSubmission(ID interface{}) (Submission, error) {
	var submission Submission
	result := DB.First(&submission, ID)
	return submission, result.Error
}
