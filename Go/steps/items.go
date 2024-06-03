package stepsPkg

import (
	"fmt"
	"time"
)

func Items(paramsChannel <-chan interface{}, responseChannel chan<- map[string]interface{}) {
	fmt.Println("Hello, items!")
	paramsToExecute := <-paramsChannel
	fmt.Println("Items: ", paramsToExecute)
	// Simular procesamiento
	time.Sleep(1000 * time.Millisecond)
	// Enviar un mensaje al canal
	responseChannel <- map[string]interface{}{"Items": []string{"item1", "item2", "item3"}}
}
