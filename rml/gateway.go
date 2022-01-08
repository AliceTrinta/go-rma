package rml

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AliceTrinta/GO-RMA/entity"
)

/*Start function will start the communication with the IoT
network and wait until receive some message*/
func Start() {
	entity.MongoConnect()
	log.Println("Mongo connected")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter JSON file content: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		//When receiving, the message will be passed to the TakeType func.
		err := TakeType(text)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(2 * time.Second)
	}
}
