package main

import (
	"GO-RMA/core/model"
	"GO-RMA/rml"
)

func main() {
	model.DeviceInstance = model.Device{}
	model.DataInstance = model.Data{}
	model.ActionInstance = model.Action{}
	rml.Start()
}

