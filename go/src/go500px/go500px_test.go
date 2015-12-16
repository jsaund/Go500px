package go500px

import "testing"

func TestGetPhotos(t *testing.T) {
	photos, err := GetPhotos("8C6ImXPi4dKEnOWC3YwPnKQO1QIYbqaystDCsijC")
	if err != nil {
		t.Fatalf("Failed to execute request. Reason: %v", err)
	}

	if photos == nil {
		t.Fatalf("Failed to get response")
	}
}
