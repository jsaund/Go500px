package api

import (
	"encoding/json"
	"model/photo"
	"net/http"
)

type getPhotosResponse struct {
	Photos []*photo.PhotoImpl `json:"photos"`
}

func GetPhotosResponse(response *http.Response) (ApiResponse, error) {
	defer response.Body.Close()

	r := &getPhotosResponse{}
	if err := json.NewDecoder(response.Body).Decode(r); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *getPhotosResponse) Get() interface{} {
	return r.Photos
}
