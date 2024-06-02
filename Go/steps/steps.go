package stepsPkg

import (
	"fmt"
	"time"
)


func MeasureTime(fn func(<-chan interface{}, chan<- map[string]interface{}), paramsChannel <-chan interface{}, responseChannel chan<- map[string]interface{}, nodeId string) {
	startTime := time.Now()
	fn(paramsChannel, responseChannel)
	endTime := time.Now()
	fmt.Println("Execution time of node: ", nodeId, endTime.Sub(startTime))
}