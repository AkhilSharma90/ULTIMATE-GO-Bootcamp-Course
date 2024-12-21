package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CatlogNodes struct {
	CatlogNodes []Catlog `json:"catlog_nodes"`
}

type Catlog struct {
	Product_id string `json: "product_id"`
	Quantity   int    `json: "quantity"`
}

func main() {
	file, _ := os.ReadFile("test.json")

	data := CatlogNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.CatlogNodes); i++ {
		fmt.Println("Product Id: ", data.CatlogNodes[i].Product_id)
		fmt.Println("Quantity: ", data.CatlogNodes[i].Quantity)
	}

}
