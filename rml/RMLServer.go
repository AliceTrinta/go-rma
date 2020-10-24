package rml

import (
	"GO-RMA/core/model"
	"GO-RMA/core/repository"
	"encoding/json"
	"log"
)

var con = repository.MongoDB{}

func Read (data string, c chan error){
	e := make(chan error)
	defer close(e)
	d := []byte(data)
	var device model.Device
	err := json.Unmarshal(d, &device)
	if err != nil{
		var data model.Data
		err = json.Unmarshal(d, &data)
		if err != nil{
			var action model.Action
			err = json.Unmarshal(d, &action)
			if err != nil {
				log.Println("Invalid Input")
				c <- err
			}
			log.Println("registering an action...")
			go delegateAction(action, e)
			c <- <- e
			return
		}
		log.Println("registering a data...")
		go saveData(data, e)
		c <- <- e
		return
	}
	log.Println("registering a device...")
	go startDevice(device, e)
	c <- <- e
	return
}

func startDevice(device model.Device, e chan error) {
	_, err := con.GetDeviceByID(device.ID)
	if err != nil{
		log.Println("Registering new Device")
		err := con.CreateDevice(device)
		if err != nil {
			log.Println("Device was not registered ")
			e <- err
			return
		}
		e <- err
		return
	}
	log.Println("This device is already registered")
	e <- nil
	return
}

func saveData(data model.Data, e chan error){
	err := con.CreateData(data)
	if err != nil {
		log.Println("It was not possible to save data ")
		e <- err
		return
	}
	log.Println("Message registered Successfully!")
	e <- err
	return

}

func delegateAction(action model.Action, e chan error) {
	err := con.CreateAction(action)
	if err != nil {
		log.Println("It was not possible to save this action ")
		e <- err
		return
	}
	log.Println("Action saved successfully!")
	device, err := con.GetDeviceByID(action.DeviceName)
	if err != nil {
		log.Println("It was not possible find the device's action ")
		e <- err
		return
	}
	content, err := json.Marshal(action)
	if err != nil{
		log.Println("It was not possible to Marshal action ")
		e <- err
		return

	}
	sendAction(device.GatewayUUID, device.UUID, content)
	return
}

func sendAction(gatewayID string, receiverID string, content []byte){
	log.Println("Delivering Message...")
	log.Println(gatewayID)
	log.Println(receiverID)
	log.Println(content)
	return
}

