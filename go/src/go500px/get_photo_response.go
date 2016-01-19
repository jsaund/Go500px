package go500px

import (
	"encoding/json"
	"io"
)

type GetPhotoResponse interface {
	GetPhoto() Photo
	GetComments() *Comments
}

type getPhotoResponse struct {
	Photo    *photo     `json:"photo"`
	Comments []*comment `json:"comments"`
	comments *Comments
}

func NewGetPhotoResponse(input io.ReadCloser) (*getPhotoResponse, error) {
	result := &getPhotoResponse{}
	if err := json.NewDecoder(input).Decode(result); err != nil {
		return nil, err
	}
	result.comments = NewComments()
	result.comments.comments = result.Comments
	return result, nil
}

func (r *getPhotoResponse) GetPhoto() Photo {
	return r.Photo
}

func (r *getPhotoResponse) GetComments() *Comments {
	return r.comments
}
