package rml

import (
	"GO-RMA/core/model"
	"testing"
)

func TestForDevice(t *testing.T){
	model.FakeConnection()
	model.DeviceInstance = model.FakeDevice{}
	err := forDevice("{\"_id\":\"device\",\"UUID\":\"8e0dae08-b360-4ace-b214-47b3d0e93f1a\",\"gatewayUUID\":\"\",\"name\":\"Farmer\",\"description\":\"Farmer's device'\",\"cycleDelayInMillis\":\"5000\",\"resourceList\":[{\"_id\":\"irrigator\",\"name\":\"irrigator 1\",\"description\":\"This is the Irrigator 1\",\"port\":\"COM3\",\"dataUnit\":\"\",\"waitTimeInMillis\":5000,\"commandList\":[{\"id\":\"ON\",\"description\":\"Irrigator is working\"},{\"id\":\"OFF\",\"description\":\"Irrigator is not working\"}]},{\"_id\":\"light\",\"name\":\"light\",\"description\":\"This is the light sensor\",\"port\":\"COM3\",\"dataUnit\":\"\",\"waitTimeInMillis\":5000,\"commandList\":[]},{\"_id\":\"humidity\",\"name\":\"humidity\",\"description\":\"This is the humidity sensor\",\"port\":\"COM3\",\"dataUnit\":\"\",\"cycleDelayInMillis\":5000,\"commandList\":[]},{\"_id\":\"temperature\",\"name\":\"temperature\",\"description\":\"This is the temperature sensor\",\"port\":\"COM3\",\"dataUnit\":\"°C\",\"waitTimeInMillis\":5000,\"commandList\":[]},{\"_id\":\"pH\",\"name\":\"pH\",\"description\":\"This is the humidity sensor\",\"port\":\"COM3\",\"dataUnit\":\"pH\",\"waitTimeInMillis\":5000,\"commandList\":[]}],\"lastUpdate\":\"2020-10-04T23:00:40.843+00:00\"}")
	if err != nil{
		t.Fatal(err)
	}
}

func TestForData(t *testing.T){
	model.FakeConnection()
	model.DataInstance = model.FakeData{}
	err := forData("{\"_id\":{\"$oid\":\"5f7a5418cdaaf862bde8b382\"},\"instant\":\"2020-10-04T23:00:38.225Z\",\"deviceName\":\"farmer\",\"resourceName\":\"humidity\",\"value\":\"dry\"}")
	if err != nil{
		t.Fatal(err)
	}
}

func TestForAction(t *testing.T) {
	model.FakeConnection()
	model.ActionInstance = model.FakeAction{}
	err := forAction("{\"_id\":\"farmer12\",\"date\":\"2020-10-04T23:00:38.225Z\",\"deviceName\":\"farmer\",\"resourceName\":\"humidity\",\"command\":\"on\"}")
	if err != nil {
		t.Fatal(err)
	}
}

