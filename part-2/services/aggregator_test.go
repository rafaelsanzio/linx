package services

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"../models"
	"../utils"
)

func TestImageExist(t *testing.T) {
	mockLine := `{"productId":"pid482","image":"https://blog.golang.org/maps"}`

	jsonLine := models.ProductJSON{}
	json.Unmarshal([]byte(mockLine), &jsonLine)

	_, err := http.Get(jsonLine.Image)
	if err != nil {
		t.Errorf("Error image does not exist")
	}
}
func TestImageIsNotExist(t *testing.T) {
	mockLine := `{"productId":"pid482","image":"http://localhost:8080/images/1.png"}`

	jsonLine := models.ProductJSON{}
	json.Unmarshal([]byte(mockLine), &jsonLine)

	_, err := http.Get(jsonLine.Image)
	if err == nil {
		t.Errorf("Error image does exist")
	}
}

func TestVerifyThreeImagesByProduct(t *testing.T) {

	mockImages := []string{
		"http://localhost:8080/images/1.png",
		"http://localhost:8080/images/2.png",
		"http://localhost:8080/images/3.png",
	}
	mockProduct := "pid123"

	mockProducts := map[string][]string{}

	mockProducts[mockProduct] = mockImages

	checkReturn := VerifyThreeImagesByProduct(mockProducts, mockProduct)
	if !checkReturn {
		t.Errorf("Error product has 3 images")
	}
}
func TestVerifyMoreThanThreeImagesByProduct(t *testing.T) {

	mockImages := []string{
		"http://localhost:8080/images/1.png",
		"http://localhost:8080/images/2.png",
		"http://localhost:8080/images/3.png",
		"http://localhost:8080/images/4.png",
	}
	mockProduct := "pid123"

	mockProducts := map[string][]string{}

	mockProducts[mockProduct] = mockImages

	checkReturn := VerifyThreeImagesByProduct(mockProducts, mockProduct)
	if !checkReturn {
		t.Errorf("Error product has 3 images")
	}
}
func TestVerifyNotThreeImagesByProduct(t *testing.T) {

	mockImages := []string{
		"http://localhost:8080/images/1.png",
		"http://localhost:8080/images/2.png",
	}
	mockProduct := "pid123"

	mockProducts := map[string][]string{}

	mockProducts[mockProduct] = mockImages

	checkReturn := VerifyThreeImagesByProduct(mockProducts, mockProduct)
	if checkReturn {
		t.Errorf("Error product has not 3 images")
	}
}

func TestAggregator(t *testing.T) {
	mockImages := []string{
		"http://localhost:8080/images/1.png",
		"http://localhost:8080/images/2.png",
		"http://localhost:8080/images/3.png",
	}
	mockProduct := "pid123"

	mockProducts := map[string][]string{}
	mockProductsToFill := map[string][]string{}

	mockProducts[mockProduct] = mockImages
	for _, value := range mockImages {

		resp, err := http.Get(value)
		if err != nil {
			continue
		}
		if resp.StatusCode != 200 {
			continue
		}

		hasThreeImages := VerifyThreeImagesByProduct(mockProducts, mockProduct)
		if !hasThreeImages && !utils.IsInArray(mockProducts[mockProduct], value) {
			mockProductsToFill[mockProduct] = append(mockProducts[mockProduct], value)
		}
	}

	mockProductsResponse := []models.Product{}
	for key, value := range mockProductsToFill {
		mockProductResponse := models.Product{
			ProductID: key,
			Images:    value,
		}

		mockProductsResponse = append(mockProductsResponse, mockProductResponse)
	}

	importProducts := Aggregator("../files/import_test.json")
	if len(importProducts) != len(mockProductsResponse) {
		t.Errorf("Error import products incorrect")
	}
}

func TestAggregatorIncorrect(t *testing.T) {
	mockImages := []string{
		"http://localhost:8080/images/1.png",
		"http://localhost:8080/images/2.png",
	}
	mockProduct := "pid123"

	mockProducts := map[string][]string{}
	mockProductsToFill := map[string][]string{}

	mockProducts[mockProduct] = mockImages
	for _, value := range mockImages {

		resp, err := http.Get(value)
		if err != nil {
			continue
		}
		if resp.StatusCode != 200 {
			continue
		}

		hasThreeImages := VerifyThreeImagesByProduct(mockProducts, mockProduct)
		if !hasThreeImages && !utils.IsInArray(mockProducts[mockProduct], value) {
			mockProductsToFill[mockProduct] = append(mockProducts[mockProduct], value)
		}
	}

	mockProductsResponse := []models.Product{}
	for key, value := range mockProductsToFill {
		mockProductResponse := models.Product{
			ProductID: key,
			Images:    value,
		}

		mockProductsResponse = append(mockProductsResponse, mockProductResponse)
	}

	importProducts := Aggregator("../files/import_image_exists_test.json")
	if len(importProducts) == len(mockProductsResponse) {
		t.Errorf("Error import products correct")
	}
}

func TestOpenFileExists(t *testing.T) {
	file, err := os.Open("../files/import_test.json")
	if err != nil {
		t.Errorf("Error could not open file")
	}
	defer file.Close()
}
func TestOpenFileDoesNotExists(t *testing.T) {
	file, err := os.Open("../files/not_exists.json")
	if err == nil {
		t.Errorf("Error could open file")
	}
	defer file.Close()
}

func TestParseCorrectLineInStruct(t *testing.T) {
	mockJSONLine := models.ProductJSON{}
	mockLine := `{"productId":"pid482","image":"http://localhost:8080/images/1.png"}`

	json.Unmarshal([]byte(mockLine), &mockJSONLine)

	if len(mockJSONLine.ProductID) == 0 {
		t.Errorf("Could not Unmarshal line into struct: %s", mockLine)
	}
}

func TestParseNotCorrectLineInStruct(t *testing.T) {
	mockJSONLine := models.ProductJSON{}
	mockLine := `{"nonProductId":"pid482","nonImage":"http://localhost:8080/images/1.png"}`

	json.Unmarshal([]byte(mockLine), &mockJSONLine)

	if len(mockJSONLine.ProductID) != 0 {
		t.Errorf("Could Unmarshal line into struct: %s", mockLine)
	}
}
