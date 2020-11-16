package rml

import (
	"GO-RMA/core/services"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func Start() {
	services.MongoConnect()
	log.Println("Mongo connected")
	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter JSON file content: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		c := make(chan error)
		go TakeType(text, c)
		var err = <-c
		if err != nil {
			log.Println(err)
		}
		close(c)
		time.Sleep(2*time.Second)
	}
}
