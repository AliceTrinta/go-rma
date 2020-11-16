package rml

import (
	"GO-RMA/core/model"
	"encoding/json"
	"github.com/tidwall/gjson"
	"log"
)

var con = model.MongoDB{}

func TakeType(input string) (err error){
	value := gjson.Get(input, "_id")
	if value.String() == "device"{
		return forDevice(input)
	} else if value.String() == "data" {
		return forData(input)
	} else if value.String() == "action" {
		return forAction(input)
	} else {
		return &model.InvalidJSONInput{}
	}
}

func forDevice(device string) (err error)  {
	log.Println("registering a device...")
	var object model.Device
	iotObject := []byte(device)
	err = json.Unmarshal(iotObject, &object)
	if err != nil {
		return
	}
	model.DeviceInstance = object
	return model.DeviceInstance.StartDevice(con)
}

func forData(data string) (err error)  {
	log.Println("registering a data...")
	var d model.Data
	dataJSON := []byte(data)
	err = json.Unmarshal(dataJSON, &d)
	if err != nil {
		return
	}
	model.DataInstance = d
	return model.DataInstance.SaveData(con)
}

func forAction(action string) (err error) {
	log.Println("registering an action...")
	actionJSON := []byte(action)
	var a model.Action
	err = json.Unmarshal(actionJSON, &a)
	if err != nil {
		return
	}
	model.ActionInstance = a
	return model.ActionInstance.DelegateAction(con)
}

