package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/AliceTrinta/GO-RMA/api/handler"
	"github.com/AliceTrinta/GO-RMA/config"
	"github.com/AliceTrinta/GO-RMA/controllers"
	"github.com/AliceTrinta/GO-RMA/infrastructure/repository"
	"github.com/AliceTrinta/GO-RMA/usecase/Data"
	"github.com/AliceTrinta/GO-RMA/usecase/Device"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {

	go controllers.RmlDevice()
	go controllers.RmlData()

	db, err := config.MongoConnect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Session.Close()

	repoDevice := repository.NewDeviceMongo(db)
	serviceDevice := Device.NewService(repoDevice)

	repoData := repository.NewDataMongo(db)
	serviceData := Data.NewService(repoData)

	r := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeDeviceHandlers(r, *n, serviceDevice)
	handler.MakeDataHandlers(r, *n, serviceData)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		time.Sleep(10 * time.Minute)
	}
}
