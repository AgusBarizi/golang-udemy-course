package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {

	jsonString := `{"id":"Z123","name":"Zenfone Max Pro","price":2500000,"image_url":"https://fdn2.gsmarena.com/vv/bigpic/asus-zenfone-max-pro-m1-zb601kl-.jpg"}`
	bytes := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
	fmt.Println(result["image_url"])

}

func TestMapEncode(t *testing.T) {

	product := map[string]interface{}{
		"id":    "P001",
		"name":  "Macbook Pro M1",
		"price": 16000000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
