package stepsPkg

import (
	"fmt"
	"time"
)

func Campaigns(paramsChannel <-chan interface{}, responseChannel chan<- map[string]interface{}) {
	fmt.Println("Hello, Campaigns!")
	paramsToExecute := <-paramsChannel
	fmt.Println("campaigns_Params: ", paramsToExecute)
	
	// Simular procesamiento
	time.Sleep(1000 * time.Millisecond)
	// Enviar un mensaje al canal
	responseChannel <- map[string]interface{}{"campaigns": []string{"campaign1", "campaign2", "campaign3"}, "originalCampaigns":[]string{"campaign1", "campaign2", "campaign3"}}
}
