package stepsPkg

import (
	"fmt"
	"time"
)

func User(paramsChannel <-chan interface{}, responseChannel chan<- map[string]interface{}) {
	fmt.Println("Hello, User!")
	paramsToExecute := <-paramsChannel
	fmt.Println("User: ", paramsToExecute)
	
	// Simular procesamiento
	time.Sleep(1000 * time.Millisecond)
	// Enviar un mensaje al canal
	responseChannel <- map[string]interface{}{"userId": 1}
}
