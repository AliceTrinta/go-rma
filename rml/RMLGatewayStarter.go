package rml

import (
	"GO-RMA/core/repository"
	"bufio"
	"fmt"
	"log"
	"os"
)

func Start() {
	repository.MongoConnect()
	log.Println("Mongo connected")
	for{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter JSON file content: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		c := make(chan error)
		go Read(text, c)
		var err = <-c
		if err != nil {
			log.Println(err)
		}
		close(c)
	}
}
