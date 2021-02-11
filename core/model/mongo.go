package model

import (
	"log"

	"github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2"
)

//MongoDB will store the database information address.
type MongoDB struct {
	Server   string
	Database string
}

//MongoRepository will gather all the functions that a MongoDB type can implement.
type MongoRepository interface {
	GetEnvironmentByUUID(UUID string) (Environment, error)
	CreateEnvironment(env Environment) error
	GetAllDevices() ([]Device, error)
	GetDeviceByUUID(UUID string) (Device, error)
	CreateDevice(device Device) error
	CreateData(data Data) error
	CreateAction(action Action) error
}

//Db variable will establish one or more connections with the cluster of servers defined by the url parameter.
var Db *mgo.Database

//ReadMongoConfig function will read the configuration file that contains the database information address.
func ReadMongoConfig() (m *MongoDB, err error) {
	if _, err = toml.DecodeFile("https://github.com/AliceTrinta/GO-RMA/blob/286e426ded75a3ab81cb6a149a3862bd94dfe6c2/core/model/mongo.toml", &m); err != nil {
		log.Fatal(err)
	}
	return
}

//MongoConnect function will establish the connection with the database.
func MongoConnect() {
	m, err := ReadMongoConfig()
	if err != nil {
		log.Fatal(err)
	}
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	Db = session.DB(m.Database)
}
