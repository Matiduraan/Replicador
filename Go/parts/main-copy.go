package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"replicador/steps"
)

type Nodes struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Params []string `json:"params"`
	Children []string `json:"children"`
}

type ProcessedNode struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Params []string `json:"params"`
	Children []string `json:"children"`
	Dependencies []string `json:"dependencies"`
}

var MOCK_STEPS = []string{"1", "2", "3","4","5"}



func main() {
	functions := make(map[string]func(<-chan interface{}, chan<- map[string]interface{}))
	functions["user"] = steps.User
	functions["items"] = steps.Items
	functions["ads"] = steps.Ads

	paramsHeap := make(map[string]interface{})

	wg := sync.WaitGroup{}


	for _, value := range functions {
		wg.Add(1)
		paramsChannel := make(chan interface{}, 1)
		resultsChannel := make(chan map[string]interface{}, 1)

		go value(paramsChannel, resultsChannel)
		paramsChannel <- "Hello"
		close(paramsChannel)
		
		go waitForResponse(resultsChannel, paramsHeap, &wg)
	}

	wg.Wait()
	fmt.Println(paramsHeap)
}

func waitForResponse(resultsChannel <-chan map[string]interface{},  paramsHeap map[string]interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Waiting for response...")
	response := <-resultsChannel
	for key, value := range response {
		paramsHeap[key] = value
	}

}




// func execNode(path string, params []interface{}) map[string]interface{} {
// 	fmt.Println("Executing node", path, "with params", params)
// }

func readJsonFile() []Node {
	jsonFile, err := os.Open("tree.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened tree.json")

	defer jsonFile.Close()

	// read our opened jsonFile and assign it to a Node array
	var nodes Nodes

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &nodes)
	return nodes.Nodes

}

func processNode(node Node, dependencyMap map[string][]string) ProcessedNode {
	var nodeId = node.Id
	var dependencies = dependencyMap[nodeId]
	return ProcessedNode{
		Id:           node.Id,
		Name:         node.Name,
		Path:         node.Path,
		Params:       node.Params,
		Children:     node.Children, // Convert node.Children to a string slice
		Dependencies: dependencies,
	}
}

func processTree(tree []Node, stepsToExecute []string) []ProcessedNode {
	var dependencyMap = buildDependencyMap(tree)
	var processedNodes []ProcessedNode = make([]ProcessedNode, len(tree))
	for _, node := range tree {
		processedNode := processNode(node, dependencyMap)
        for _, step := range stepsToExecute {
            if step == processedNode.Id {
                processedNodes = append(processedNodes, processedNode)
                break
            }
        }
	}
	
	return processedNodes
}

func buildDependencyMap(tree []Node) map[string][]string {
    dependencyMap := make(map[string][]string)

    for _, node := range tree {
		for _, childID := range node.Children {
			strChildID := string(childID)
            if _, exists := dependencyMap[strChildID]; !exists {
                dependencyMap[strChildID] = []string{}
            }
            dependencyMap[strChildID] = append(dependencyMap[strChildID], node.Id)
        }
    }

    return dependencyMap
}

func findNodesToExecute(finishedNodes []ProcessedNode, processedNodes []ProcessedNode) []ProcessedNode {
	finishedNodeIds := make(map[string]bool)
    for _, node := range finishedNodes {
        finishedNodeIds[node.Id] = true
    }

    var nodesToExecute []ProcessedNode
    for _, node := range processedNodes {
        if !finishedNodeIds[node.Id] {
            allDependenciesFinished := true
            for _, dependency := range node.Dependencies {
                if !finishedNodeIds[dependency] {
                    allDependenciesFinished = false
                    break
                }
            }
            if allDependenciesFinished {
                nodesToExecute = append(nodesToExecute, node)
            }
        }
    }

    return nodesToExecute
}