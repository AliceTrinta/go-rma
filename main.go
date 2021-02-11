package main

import (
	"github.com/AliceTrinta/GO-RMA/core/model"
	"github.com/AliceTrinta/GO-RMA/rml"
)

func main() {
	model.DeviceInstance = model.Device{}
	model.DataInstance = model.Data{}
	model.ActionInstance = model.Action{}
	rml.Start()
}
