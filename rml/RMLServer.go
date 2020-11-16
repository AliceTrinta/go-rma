package rml

import (
	"GO-RMA/core/model"
	"GO-RMA/core/services"
	"encoding/json"
	"github.com/tidwall/gjson"
	"log"
)

var con = services.MongoDB{}

func TakeType(input string, c chan error){
	e := make(chan error)
	value := gjson.Get(input, "_id")
	if value.String() == "device"{
		forDevice(input, e)
		c <- <- e
		return
	} else if value.String() == "data" {
		forData(input, e)
		c <- <- e
		return
	} else if value.String() == "action" {
		forAction(input, e)
		c <- <- e
		return
	} else {
		c <- &model.InvalidJSONInput{}
	}
}

func forDevice(device string, e chan error)  {
	log.Println("registering a device...")
	c := make(chan error)
	var object services.Device
	iotObject := []byte(device)
	err := json.Unmarshal(iotObject, &object)
	if err != nil {
		e <- err
		return
	}
	object.StartDevice(c, con)
	e <- <- c
}

func forData(data string, e chan error)  {
	log.Println("registering a data...")
	c := make(chan error)
	var iotData services.Data
	dataJSON := []byte(data)
	err := json.Unmarshal(dataJSON, &iotData)
	if err != nil {
		e <- err
		return
	}
	iotData.SaveData(c, con)
	e <- <- c

}

func forAction(action string, e chan error)  {
	log.Println("registering an action...")
	c := make(chan error)
	var iotAction services.Action
	actionJSON := []byte(action)
	err := json.Unmarshal(actionJSON, &iotAction)
	if err != nil {
		e <- err
		return
	}
	iotAction.DelegateAction(c, con)
	e <- <- c
}

