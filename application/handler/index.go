package handler

import (
	"net/http"
)

//Index function will be called by the index page.
func Index(w http.ResponseWriter, r *http.Request) {
	//var con = model.MongoDB{}
	//devices, err := con.GetAllDevices()
	// if err != nil {
	// 	log.Println("It was not possible to get Devices.")
	// }
	msg := []byte("Welcome to Wonder")
	w.Write(msg)
}
