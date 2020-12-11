package model

import (
	"github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2"
	"log"
)

//The MongoDB will store the database information address.
type MongoDB struct {
	Server string
	Database string
}

//The MongoRepository will gather all the functions that a MongoDB type can implement.
type MongoRepository interface {
	GetEnvironmentByUUID(UUID string) (Environment, error)
	CreateEnvironment(env Environment) error
	GetDeviceByUUID(UUID string) (Device, error)
	CreateDevice(device Device) error
	CreateData(data Data) error
	CreateAction(action Action) error
}

//The Db variable will establish one or more connections with the cluster of servers defined by the url parameter.
var Db *mgo.Database

//The ReadMongoConfig function will read the configuration file that contains the database information address.
func ReadMongoConfig() (m *MongoDB, err error) {
	if _, err = toml.DecodeFile("C:\\Users\\ATG20\\go\\src\\GO-RMA\\core\\model\\mongo.toml", &m); err != nil {
		log.Fatal(err)
	}
	return
}

//The MongoConnect function will establish the connection with the database.
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
