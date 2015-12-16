package go500px

import (
	"testing"
)

func TestSimpleRequest(t *testing.T) {
	manager := NewApiManager("8C6ImXPi4dKEnOWC3YwPnKQO1QIYbqaystDCsijC")

	getPhotosRequest, err := GetPhotosBuilder().
		Feature("fresh_today").
		Sort("highest_rating").
		Build()

	if err != nil {
		t.Fatalf("Failed to create getPhotosApiRequest. Reason: %v", err)
	}

	apiResponse, err := manager.Do(getPhotosRequest)

	if err != nil {
		t.Fatalf("Failed to execute getPhotosApiRequest. Reason: %v", err)
	}

	if apiResponse == nil {
		t.Fatalf("Failed to execute getPhotosApiRequest. Reason: api response is nil")
	}

	if photos, ok := apiResponse.response.([]*photoImpl); !ok {
		t.Fatalf("Expected a slice of photos to be returned.")
	} else {
		t.Logf("Name: %s", photos[0].Name())
	}
}
