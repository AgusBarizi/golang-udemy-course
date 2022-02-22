package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:      "Monkey",
		MiddleName:     "D",
		LastName:       "Luffy",
		Age:            0,
		Weight:         0,
		HaveVaccinated: false,
		Hobbies:        []string{"Game", "Badmiton", "Archery", "Coding", "Cycling"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Monkey","MiddleName":"D","LastName":"Luffy","Age":0,"Weight":0,"HaveVaccinated":false,"Hobbies":["Game","Badmiton","Archery","Coding","Cycling"]}`
	jsonBytes := []byte(jsonString)

	var customer = &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:      "Monkey",
		MiddleName:     "D",
		LastName:       "Luffy",
		Age:            22,
		Weight:         50.6,
		HaveVaccinated: false,
		Hobbies:        []string{"Minum Sake"},
		Addresses: []Address{
			{
				City:       "Jakarta",
				PostalCode: 123,
				IsMain:     false,
			},
			{
				City:       "Surabaya",
				PostalCode: 111,
				IsMain:     true,
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Monkey","MiddleName":"D","LastName":"Luffy","Age":22,"Weight":50.6,"HaveVaccinated":false,"Hobbies":["Minum Sake"],"Addresses":[{"City":"Jakarta","PostalCode":123,"IsMain":false},{"City":"Surabaya","PostalCode":111,"IsMain":true}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"City":"Jakarta","PostalCode":123,"IsMain":false},{"City":"Surabaya","PostalCode":111,"IsMain":true}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			City:       "Jakarta",
			PostalCode: 123,
			IsMain:     false,
		},
		{
			City:       "Surabaya",
			PostalCode: 111,
			IsMain:     true,
		},
	}

	bytes, err := json.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
