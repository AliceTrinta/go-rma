package main

import (
	"time"

	"github.com/AliceTrinta/GO-RMA/controllers"
)

func main() {

	go controllers.RmlDevice()
	go controllers.RmlData()

	for {
		time.Sleep(10*time.Minute)
	}
}
