package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/AliceTrinta/GO-RMA/usecase/Device"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func ListDevices(service Device.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading devices"
		devices, err := service.GetAllDevices()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*entity.Device
		for _, d := range devices {
			toJ = append(toJ, &entity.Device{
				UUID:               d.UUID,
				GatewayUUID:        d.GatewayUUID,
				Name:               d.Name,
				Description:        d.Description,
				CycleDelayInMillis: d.CycleDelayInMillis,
				ResourceList:       d.ResourceList,
				Status:             d.Status,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func MakeDeviceHandlers(r *mux.Router, n negroni.Negroni, service Device.UseCase) {
	r.Handle("/listDevices", n.With(
		negroni.Wrap(ListDevices(service)),
	)).Methods("GET", "OPTIONS").Name("ListDevices")
}
