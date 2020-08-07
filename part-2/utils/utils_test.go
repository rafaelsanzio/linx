package utils

import (
	"testing"
)

func TestIsInArray(t *testing.T) {
	mockImages := []string{
		"http://localhost:8080/images/1.png",
		"http://localhost:8080/images/2.png",
	}

	mockImage := "http://localhost:8080/images/1.png"

	verify := IsInArray(mockImages, mockImage)
	if !verify {
		t.Errorf("Error image has in array")
	}
}
func TestIsNotInArray(t *testing.T) {
	mockImages := []string{
		"http://localhost:8080/images/1.png",
		"http://localhost:8080/images/2.png",
	}

	mockImage := "http://localhost:8080/images/3.png"

	verify := IsInArray(mockImages, mockImage)
	if verify {
		t.Errorf("Error image has not in array")
	}
}
