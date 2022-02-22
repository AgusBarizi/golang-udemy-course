package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	City       string
	PostalCode int64
	IsMain     bool
}

type Customer struct {
	FirstName      string
	MiddleName     string
	LastName       string
	Age            int64
	Weight         float64
	HaveVaccinated bool
	Hobbies        []string
	Addresses      []Address
}

func TestJSONObject(t *testing.T) {

	customer := Customer{
		FirstName:      "Monkey",
		MiddleName:     "D",
		LastName:       "Luffy",
		Age:            25,
		Weight:         49.90,
		HaveVaccinated: true,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
