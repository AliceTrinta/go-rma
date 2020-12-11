package rml

import (
	"GO-RMA/core/model"
	"testing"
)

//Testing the forDevice func.
func TestForDevice(t *testing.T){
	model.FakeConnection()
	model.DeviceInstance = model.FakeDevice{}
	err := forDevice("{\"UUID\":\"\",\"gatewayUUID\":\"\",\"name\":\"\",\"description\":\"\",\"cycleDelayInMillis\":\"\",\"resourceList\":[],\"status\":false}")
	if err != nil{
		t.Fatal(err)
	}
}

//Testing the forData func.
func TestForData(t *testing.T){
	model.FakeConnection()
	model.DataInstance = model.FakeData{}
	err := forData("{\"_id\":1,\"instant\":\"0001-01-01T00:00:00.000Z\",\"UUID\":\"\",\"resourceName\":\"\",\"value\":\"\"}")
	if err != nil{
		t.Fatal(err)
	}
}

//Testing the forAction func.
func TestForAction(t *testing.T) {
	model.FakeConnection()
	model.ActionInstance = model.FakeAction{}
	err := forAction("{\"_id\":1,\"date\":\"0001-01-01T00:00:00.000Z\",\"UUID\":\"\",\"resourceName\":\"\",\"command\":\"\"}")
	if err != nil {
		t.Fatal(err)
	}
}

