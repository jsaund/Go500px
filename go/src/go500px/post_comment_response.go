package go500px

import (
	"encoding/json"
	"io"
)

type PostCommentResponse interface {
	GetStatus() string
	GetMessage() string
	GetError() string
}

type postCommentResponse struct {
	CommentStatus  string `json:"status"`
	CommentMessage string `json:"message"`
	CommentError   string `json:"error"`
}

func NewPostCommentResponse(input io.ReadCloser) (*postCommentResponse, error) {
	result := &postCommentResponse{}
	if err := json.NewDecoder(input).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *postCommentResponse) GetStatus() string {
	return r.CommentStatus
}

func (r *postCommentResponse) GetMessage() string {
	return r.CommentMessage
}

func (r *postCommentResponse) GetError() string {
	return r.CommentError
}
