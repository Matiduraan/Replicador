package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
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

var MOCK_STEPS = []string{"ads", "items", "user", "userPreferences", "campaigns"}

func main() {
	fmt.Println("Hello, World!")
	tree := readJsonFile()
	processedTree := processTree(tree, MOCK_STEPS)
	var finishedNodes []ProcessedNode
	paramsHeap := make(map[string]interface{})
	functions := GetNeededFunctions(MOCK_STEPS)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for {
		nodesToExecute := findNodesToExecute(finishedNodes, processedTree)
		fmt.Println("Nodes to execute", nodesToExecute)
		if len(finishedNodes) == len(processedTree) {
			break
		}
		for _, node := range nodesToExecute {
			wg.Add(1)
			go func(node ProcessedNode) {
				defer wg.Done()
				fmt.Println("Executing node", node.Name)
				var params = make(map[string]interface{})
				for _, param := range node.Params {
					mu.Lock()
					if paramsHeap[param] == nil {
						mu.Unlock()
						panic("Missing parameter " + param)
					}
					params[param] = paramsHeap[param]
					mu.Unlock()
				}
				paramsChannel := make(chan interface{}, 1)
				resultsChannel := make(chan map[string]interface{}, 1)
				go functions[node.Name](paramsChannel, resultsChannel)
				paramsChannel <- params
				close(paramsChannel)
				response := <-resultsChannel
				for key, value := range response {
					paramsHeap[key] = value
				}
				mu.Lock()
				finishedNodes = append(finishedNodes, node)
				mu.Unlock()
				fmt.Println("Finished executing node", node.Name)
			}(node)
		}
		wg.Wait()
	}
	fmt.Println("Finished executing all nodes")
}

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
	var processedNodes []ProcessedNode = make([]ProcessedNode,0)
	for _, node := range tree {
		processedNode := processNode(node, dependencyMap)
        for _, step := range stepsToExecute {
            if step == processedNode.Name {
				fmt.Println("Processing node", processedNode)
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