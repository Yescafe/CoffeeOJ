package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	CreatorID uint   `gorm:"foreignKey:UserId"`
	Title     string `gorm:"not null;size:200"`
	MemoLimit int64  `gorm:"not null"`
	TimeLimit int64  `gorm:"not null"`
	Path      string `gorm:"size:1000"`
}

// GetProblem 用 ID 获取题目
func GetProblem(ID interface{}) (Problem, error) {
	var problem Problem
	result := DB.First(&problem, ID)
	return problem, result.Error
}
