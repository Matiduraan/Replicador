package steps

import (
	"fmt"
	"time"
)

func UserPreferences(paramsChannel <-chan interface{}, responseChannel chan<- map[string]interface{})  {
	fmt.Println("Hello, UserPreferences!")
	paramsToExecute := <-paramsChannel
	fmt.Println("preferences_Params: ", paramsToExecute)
	
	// Simular procesamiento
	time.Sleep(1000 * time.Millisecond)
	// Enviar un mensaje al canal
	responseChannel <- make(map[string]interface{})
}
