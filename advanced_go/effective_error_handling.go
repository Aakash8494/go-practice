package main

import (
	"errors"
	"fmt"
)

type Truck struct {
	id string
}

// processTruck handles the loading and unloading of a truck.
// Now it returns an 'error' type.
func processTruck(truck Truck) error {
	// Let's simulate an error for Truck-2 just to test it
	if truck.id == "Truck-2" {
		return errors.New("cargo door jammed")
	}

	fmt.Printf("Processing truck: %s\n", truck.id)
	return nil // nil means no error happened
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)

		// Catch the error returned by the function
		err := processTruck(truck)

		// Check if there is an actual error
		if err != nil {
			fmt.Printf(
				"Error: Could not process %s because %v\n",
				truck.id,
				err,
			) // here printf is used to format the error message with the truck id and the error details
			// log.Printf("Error processing truck: %s", err) // This is a more standard way to log errors in Go, it will include the timestamp and other details.
			// log.Fatalf("Error processing truck: %s", err) // This will log the error and then exit the program immediately, which is not ideal in this case since we want to continue processing other trucks.
			// fmt.Printf("Error processing truck: %s\n", err) // This is a simple way to print the error message to the console, but it won't include any additional context or formatting.
			// fmt.Println("Error processing truck: ", err) // This is another way to print the error message, it will concatenate the string with the error message, but it won't format it nicely.
			// fmt.Printf("Error processing truck: %s\n", err.Error()) // This will call the Error() method on the error object, which returns the error message as a string, but it's essentially the same as just using %s with the error object directly.
			// fmt.Printf("Error processing truck: %v\n", err) // This will use the default format for the error object, which is usually the same as %s, but it can be useful if the error type has a custom String() method that provides more details.
			// fmt.Printf("Error processing truck: %s\n", err.Error()) // This is redundant because using %s with the error object will automatically call the Error() method, so it's better to just use %s directly with the error object.
		} else {
			fmt.Println("Processed successfully.")
		}
		fmt.Println("---")
	}
	fmt.Println("-----------------------")
	fmt.Println("-----------------------")

	for _, truck := range trucks {
		if err := processTruck(truck); err != nil {
			// log.Fatalf("Error processing truck: %s", err)
			fmt.Println("Error processing truck: Println --->", err)
			fmt.Printf("Error processing truck: Printf ---> %s\n", err)
		}
	}
}
