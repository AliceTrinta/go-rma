package services

import (
	"GO-RMA/core/model"
	"encoding/json"
	"log"
	"time"
)

type Device struct {
	ID string                      `json:"_id" bson:"_id"`
	UUID string                    `json:"UUID" bson:"UUID"`
	GatewayUUID string             `json:"gateway_UUID" bson:"gatewayUUID"`
	Name string                    `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	CycleDelayInMillis json.Number `json:"cycleDelayInMillis" bson:"cycleDelayInMillis"`
	ResourceList []model.Resource  `json:"resourceList" bson:"resourceList"`
	LastUpdate time.Time           `json:"lastUpdate" bson:"lastUpdate"`
}

type IoTObject interface {
	StartDevice(chan error)
}

func (d Device)StartDevice(e chan error, con MongoDB){
	_, err := con.GetDeviceByID(d.ID)
	if err != nil{
		log.Println("Registering new Device")
		err := con.CreateDevice(d)
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

func (m *MongoDB) GetDeviceByID(id string) (Device, error) {
	var device Device
	err := Db.C("device").FindId(id).One(&device)
	return device, err
}

func (m *MongoDB) CreateDevice(device Device) error {
	err := Db.C("device").Insert(&device)
	return err
}