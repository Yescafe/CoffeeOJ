package serializer

import "singo/model"

type Submission struct {
	ID         uint   `json:"id"`
	ProblemID  uint   `json:"problem_id"`
	UserID     uint   `json:"user_id"`
	SourcePath string `json:"source_path"`
	CreatedAt  int64  `json:"created_at"`
}

func BuildSubmission(submission model.Submission) Submission {
	return Submission{
		ID:         submission.ID,
		ProblemID:  submission.ProblemID,
		UserID:     submission.UserID,
		SourcePath: submission.SourcePath,
		CreatedAt:  submission.CreatedAt.Unix(),
	}
}

func BuildSubmissionResponse(submission model.Submission) Response {
	return Response{
		Data: BuildSubmission(submission),
	}
}
