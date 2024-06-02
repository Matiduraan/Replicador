package main

import (
	"fmt"
	"replicador/steps"
)

type stepFunction func(<-chan interface{}, chan<- map[string]interface{})

func FunctionsMap() map[string]stepFunction{
	// Define a map with the function names as keys and the function pointers as values
	steps := map[string]stepFunction{
		"items":            steps.Items,
		"ads":              steps.Ads,
		"user":             steps.User,
		"userPreferences":  steps.UserPreferences,
		"campaigns":        steps.Campaigns,
	}

	// return the map with the function names as keys and the function pointers as values
	return steps
}

func GetNeededFunctions(functions []string) map[string]stepFunction{
	// Define a map with the function names as keys and the function pointers as values
	steps := FunctionsMap()
	fmt.Println("Steps: ", steps)
	neededFunctions := make(map[string]stepFunction)

	// Iterate over the functions array
	for _, function := range functions {
		// If the function exists in the map
		if _, ok := steps[function]; ok {
			// Add the function to the neededFunctions map
			neededFunctions[function] = steps[function]
		}
	}

	// return the neededFunctions map
	return neededFunctions
}
