package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/AliceTrinta/GO-RMA/usecase/Data"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func ListData(service Data.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading data"
		data, err := service.GetAllData()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*entity.Data
		for _, d := range data {
			toJ = append(toJ, &entity.Data{
				Instant:      d.Instant,
				DeviceUUID:   d.DeviceUUID,
				ResourceName: d.ResourceName,
				Value:        d.Value,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func MakeDataHandlers(r *mux.Router, n negroni.Negroni, service Data.UseCase) {
	r.Handle("/listData", n.With(
		negroni.Wrap(ListData(service)),
	)).Methods("GET", "OPTIONS").Name("ListData")
}
