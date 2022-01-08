package rml

import (
	"encoding/json"
	"log"

	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/tidwall/gjson"
)

//Creating a variable to connect with the database.
var con = entity.MongoDB{}

/*TakeType function will recognize if the message
received must be treated as a Device, a Data or an Action. */
func TakeType(input string) (err error) {
	value := gjson.Get(input, "_id")
	if value.String() == "device" {
		return forDevice(input)
	} else if value.String() == "data" {
		return forData(input)
	} else if value.String() == "action" {
		return forAction(input)
	} else {
		return &entity.InvalidJSONInput{}
	}
}

//The forDevice function will treat a Device JSON string.
func forDevice(device string) (err error) {
	log.Println("registering a device...")
	var object entity.Device
	iotObject := []byte(device)
	err = json.Unmarshal(iotObject, &object)
	if err != nil {
		return
	}
	entity.DeviceInstance = object
	return entity.DeviceInstance.StartDevice(con)
}

//The forData function will treat a Data JSON string.
func forData(data string) (err error) {
	log.Println("registering a data...")
	var d entity.Data
	dataJSON := []byte(data)
	err = json.Unmarshal(dataJSON, &d)
	if err != nil {
		return
	}
	entity.DataInstance = d
	return entity.DataInstance.SaveData(con)
}

//The forAction function will treat an Action JSON string.
func forAction(action string) (err error) {
	log.Println("registering an action...")
	actionJSON := []byte(action)
	var a entity.Action
	err = json.Unmarshal(actionJSON, &a)
	if err != nil {
		return
	}
	entity.ActionInstance = a
	return entity.ActionInstance.DelegateAction(con)
}
