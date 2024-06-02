package main

import (
	"fmt"
	"replicador/steps"
)

type stepFunction func(<-chan interface{}, chan<- map[string]interface{})

func FunctionsMap() map[string]stepFunction{
	steps := map[string]stepFunction{
		"items":            steps.Items,
		"ads":              steps.Ads,
		"user":             steps.User,
		"userPreferences":  steps.UserPreferences,
		"campaigns":        steps.Campaigns,
	}

	return steps
}

func GetNeededFunctions(functions []string) map[string]stepFunction{
	steps := FunctionsMap()
	fmt.Println("Steps: ", steps)
	neededFunctions := make(map[string]stepFunction)

	for _, function := range functions {
		if _, ok := steps[function]; ok {
			neededFunctions[function] = steps[function]
		}
	}

	return neededFunctions
}
