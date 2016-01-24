package go500px

import (
	"encoding/json"
	"io"
	"strconv"
)

type PostCommentResponse interface {
	GetStatus() string
	GetMessage() string
	GetError() string
}

type postCommentResponse struct {
	CommentStatus  int    `json:"status"`
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
	return strconv.Itoa(r.CommentStatus)
}

func (r *postCommentResponse) GetMessage() string {
	return r.CommentMessage
}

func (r *postCommentResponse) GetError() string {
	return r.CommentError
}
