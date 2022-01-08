package main

import (
	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/AliceTrinta/GO-RMA/rml"
)

func main() {
	entity.DeviceInstance = entity.Device{}
	entity.DataInstance = entity.Data{}
	entity.ActionInstance = entity.Action{}
	rml.Start()
}
