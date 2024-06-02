package main

import (
	"fmt"
	stepsPkg "replicador/steps"
)

type StepFunction func(<-chan interface{}, chan<- map[string]interface{})

func FunctionsMap() map[string]StepFunction{
	steps := map[string]StepFunction{
		"items":            stepsPkg.Items,
		"ads":              stepsPkg.Ads,
		"user":             stepsPkg.User,
		"userPreferences":  stepsPkg.UserPreferences,
		"campaigns":        stepsPkg.Campaigns,
	}

	return steps
}

func GetNeededFunctions(functions []string) map[string]StepFunction{
	steps := FunctionsMap()
	fmt.Println("Steps: ", steps)
	neededFunctions := make(map[string]StepFunction)

	for _, function := range functions {
		if _, ok := steps[function]; ok {
			neededFunctions[function] = steps[function]
		}
	}

	return neededFunctions
}
