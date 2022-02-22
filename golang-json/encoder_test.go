package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, err := os.Create("customer_output.json")
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(writer)
	customer := &Customer{
		FirstName:      "Monkey",
		MiddleName:     "D",
		LastName:       "Luffy",
		Age:            35,
		Weight:         65.9,
		HaveVaccinated: true,
		Hobbies:        []string{"Football", "Game"},
		Addresses:      nil,
	}
	err = encoder.Encode(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
