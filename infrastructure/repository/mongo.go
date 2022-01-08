package repository

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

// //MongoRepository will gather all the functions that a MongoDB type can implement.
// type MongoRepository interface {
// 	GetEnvironmentByUUID(UUID string) (entity.Environment, error)
// 	CreateEnvironment(env entity.Environment) error
// 	GetAllDevices() ([]entity.Device, error)
// 	GetDeviceByUUID(UUID string) (entity.Device, error)
// 	CreateDevice(device entity.Device) error
// 	CreateData(data entity.Data) error
// 	CreateAction(action entity.Action) error
// }

//Db variable will establish one or more connections with the cluster of servers defined by the url parameter.
var Db *mgo.Database

//ReadMongoConfig function will read the configuration file that contains the database information address.
func ReadMongoConfig() (m *MongoDB, err error) {
	if _, err = toml.DecodeFile("/home/trinta/go/src/github.com/AliceTrinta/GO-RMA/infrastructure/repository/mongo.toml", &m); err != nil {
		log.Fatal(err)
	}
	return
}

//MongoConnect function will establish the connection with the database.
func MongoConnect() *mgo.Database {
	m, err := ReadMongoConfig()
	if err != nil {
		log.Fatal(err)
	}
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	return session.DB(m.Database)
}
