package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"./services"
)

func main() {
	productsResponse := services.Aggregator(os.Args[1])

	dumpFile, _ := json.MarshalIndent(productsResponse, "", " ")
	_ = ioutil.WriteFile("dump.json", dumpFile, 0644)
}
