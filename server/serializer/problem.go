package serializer

import "singo/model"

type Problem struct {
	ID        uint   `json:"id"`
	CreatorID uint   `json:"creator_id"`
	Title     string `json:"title"`
	MemoLimit int64  `json:"memo_limit"`
	TimeLimit int64  `json:"time_limit""`
	Path      string `json:"path"`
	CreatedAt int64  `json:"created_at"`
}

func BuildProblem(problem model.Problem) Problem {
	return Problem{
		ID:        problem.ID,
		CreatorID: problem.CreatorID,
		Title:     problem.Title,
		MemoLimit: problem.MemoLimit,
		TimeLimit: problem.TimeLimit,
		Path:      problem.Path,
		CreatedAt: problem.CreatedAt.Unix(),
	}
}

func BuildProblemResponse(problem model.Problem) Response {
	return Response{
		Data: BuildProblem(problem),
	}
}
