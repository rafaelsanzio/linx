package services

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"../models"
	"../utils"
)

// Aggregator service to aggregate images products
func Aggregator(filename string) []models.Product {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	jsonLine := models.ProductJSON{}
	products := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		err := json.Unmarshal([]byte(line), &jsonLine)
		if err != nil {
			log.Println("Could not Unmarshal line into struct: ", line)
			continue
		}

		resp, err := http.Get(jsonLine.Image)
		if err != nil || resp.StatusCode != 200 {
			continue
		}

		hasThreeImages := VerifyThreeImagesByProduct(products, jsonLine.ProductID)
		if !hasThreeImages && !utils.IsInArray(products[jsonLine.ProductID], jsonLine.Image) {
			products[jsonLine.ProductID] = append(products[jsonLine.ProductID], jsonLine.Image)
		}
	}

	productsResponse := []models.Product{}
	for key, value := range products {
		product := models.Product{
			ProductID: key,
			Images:    value,
		}

		productsResponse = append(productsResponse, product)
	}

	return productsResponse
}

// VerifyThreeImagesByProduct with products has three images already filled
func VerifyThreeImagesByProduct(products map[string][]string, productID string) bool {
	if len(products[productID]) >= 3 {
		return true
	}
	return false
}
