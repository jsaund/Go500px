package go500px

import "testing"

func TestGetPhotos(t *testing.T) {
	s := GetPhotos("8C6ImXPi4dKEnOWC3YwPnKQO1QIYbqaystDCsijC")
	if s == "" {
		t.Fatalf("Failed to get response")
	}
	t.Logf("Response: %s\n", s)
}
