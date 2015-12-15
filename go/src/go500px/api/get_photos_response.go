package api

import (
	"encoding/json"
	"go500px/model/photo"
	"net/http"
)

type getPhotosResponse struct {
	Photos []*photo.PhotoImpl `json:"photos"`
}

func GetPhotosResponse(response *http.Response) (*ApiResponse, error) {
	defer response.Body.Close()

	r := &getPhotosResponse{}
	if err := json.NewDecoder(response.Body).Decode(r); err != nil {
		return nil, err
	}

	return &ApiResponse{r.Photos}, nil
}
