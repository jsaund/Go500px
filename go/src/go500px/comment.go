package go500px

import "time"

type Comment interface {
	ID() int
	UserID() int
	ToWhomUserID() int
	Body() string
	CreatedAt() int64
	ParentID() int
	User() User
}

type Comments struct {
	comments []*comment
}

func NewComments() *Comments {
	return new(Comments)
}

func (c *Comments) Count() int32 {
	return int32(len(c.comments))
}

func (c *Comments) Get(index int32) Comment {
	return c.comments[index]
}

type comment struct {
	CommentID           int        `json:"id"`
	CommentUserID       int        `json:"user_id"`
	CommentToWhomUserID int        `json:"to_whom_user_id"`
	CommentBody         string     `json:"body"`
	CommentCreatedAt    *time.Time `json:"created_at"`
	CommentParentID     int        `json:"parent_id"`
	CommentUser         *user      `json:"user"`
}

func (c *comment) ID() int {
	return c.CommentID
}

func (c *comment) UserID() int {
	return c.CommentUserID
}

func (c *comment) ToWhomUserID() int {
	return c.CommentToWhomUserID
}

func (c *comment) Body() string {
	return c.CommentBody
}

func (c *comment) CreatedAt() int64 {
	return int64(c.CommentCreatedAt.Nanosecond() / 1000)
}

func (c *comment) ParentID() int {
	return c.CommentParentID
}

func (c *comment) User() User {
	return c.CommentUser
}
