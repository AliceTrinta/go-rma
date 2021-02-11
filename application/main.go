package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AliceTrinta/GO-RMA/application/controller"
	"github.com/AliceTrinta/GO-RMA/core/model"
)

func init() {
	//Connect to your MongoDatabase
	model.MongoConnect()
	fmt.Println("Mongo connected")
}

func main() {
	handler := controller.Begin()
	svr := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      handler,
	}
	log.Fatal(svr.ListenAndServe())
}
