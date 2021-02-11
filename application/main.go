package main

import (
	"GO-RMA/application/controller"
	"GO-RMA/core/model"
	"fmt"
	"log"
	"net/http"
	"time"
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
