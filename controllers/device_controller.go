package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/AliceTrinta/GO-RMA/infrastructure/repository"
	"github.com/AliceTrinta/GO-RMA/usecase/Device"
	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
)

func RmlDevice() error {
	nc, err := nats.Connect("nats://demo.nats.io:4222")
	if err != nil {
		log.Fatal(err)
	}
	sub, err := nc.SubscribeSync("connect")
		if err != nil {
			log.Fatal(err)
		}

	for {
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			continue
		}
		uuid, err := StartDevice(msg.Data)
		if err == nil {
			msg.Respond([]byte(uuid))
		}
	}
}

func StartDevice(device []byte) (string, error) {
	db := repository.MongoConnect()
	defer db.Session.Close()
	repo := repository.NewDeviceMongo(db)
	service := Device.NewService(repo)

	var object entity.Device
	err := json.Unmarshal(device, &object)
	if err != nil {
		return "", err
	}

	_, err = service.GetDeviceByUUID(object.UUID)
	if err == nil {
		log.Println("Device already exists")
		return "", err
	}

	log.Println("registering device...")
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err.Error())
	}
	object.UUID = uuid.String()
	err = service.CreateDevice(object)
	if err != nil {
		log.Fatal("error creating device: " + err.Error())
		return "", err
	}

	return object.UUID, nil

}
