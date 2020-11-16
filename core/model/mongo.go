package model

import (
	"github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2"
	"log"
)

type MongoDB struct {
	Server string
	Database string
}

type MongoRepository interface {
	GetDeviceByID(id string) (Device, error)
	CreateDevice(device Device) error
	CreateData(data Data) error
	CreateAction(action Action) error
}

var Db *mgo.Database

func ReadMongoConfig() (m *MongoDB) {
	if _, err := toml.DecodeFile("C:\\Users\\ATG20\\go\\src\\GO-RMA\\core\\model\\mongo.toml", &m); err != nil {
		log.Fatal(err)
	}
	return
}

func MongoConnect() {
	var m = ReadMongoConfig()
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	Db = session.DB(m.Database)
}
