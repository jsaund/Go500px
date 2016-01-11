package go500px

import (
	"encoding/json"
	"io"
)

type GetPhotosResponse interface {
	GetPhotos() *Photos
}

type getPhotosResponse struct {
	Photos []*photo `json:"photos"`
	photos *Photos
}

func NewGetPhotosResponse(input io.ReadCloser) (*getPhotosResponse, error) {
	result := &getPhotosResponse{}
	if err := json.NewDecoder(input).Decode(result); err != nil {
		return nil, err
	}
	result.photos = NewPhotos()
	result.photos.photos = result.Photos
	return result, nil
}

func (r *getPhotosResponse) GetPhotos() *Photos {
	return r.photos
}
