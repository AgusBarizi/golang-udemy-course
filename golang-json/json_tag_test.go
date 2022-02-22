package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "Z123",
		Name:     "Zenfone Max Pro",
		ImageURL: "https://fdn2.gsmarena.com/vv/bigpic/asus-zenfone-max-pro-m1-zb601kl-.jpg",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"Z123","name":"Zenfone Max Pro","image_url":"https://fdn2.gsmarena.com/vv/bigpic/asus-zenfone-max-pro-m1-zb601kl-.jpg"}`
	bytes := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(bytes, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageURL)
}
