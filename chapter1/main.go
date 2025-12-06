package main

import (
	"fmt"
)

type Truck struct {
	id string
}

//processTruck handles the loading and unloading of a truck

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck with ID: %s\n", truck.id)
	return nil
}

func main() {
	trucks := []Truck{
		{id: "TRK123"},
		{id: "TRK456"},
		{id: "TRK789"},
	}

	for _, truck := range trucks {
		if err := processTruck(truck); err != nil {
			fmt.Printf("Error processing truck %s: %v\n", truck.id, err)
		}
	}
}
