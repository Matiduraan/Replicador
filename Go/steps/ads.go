package steps

import (
	"fmt"
	"time"
)

func Ads(paramsChannel <-chan interface{}, responseChannel chan<- map[string]interface{}) {
	fmt.Println("Hello, Ads!")
	paramsToExecute := <-paramsChannel
	fmt.Println("Ads: ", paramsToExecute)
	// Simular procesamiento
	time.Sleep(1000 * time.Millisecond)
	// Enviar un mensaje al canal
	responseChannel <- map[string]interface{}{"Ads": []string{"ad1", "ad2", "ad3"}}
}
