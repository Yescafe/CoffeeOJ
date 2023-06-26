package service

import (
	"singo/model"
	"singo/serializer"
)

type ProblemListService struct{}

func (service *ProblemListService) List() serializer.Response {
	var problems []model.Problem
	model.DB.Find(&problems)

	var problemIDs []uint
	for i := range problems {
		problemIDs = append(problemIDs, problems[i].ID)
	}

	return serializer.Response{
		Code: 200,
		Data: problemIDs,
	}
}
