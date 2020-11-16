package services

import (
	"github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2"
	"log"
)

type MongoDB struct {
	Server string
	Database string
}

var Db *mgo.Database

func readMongoConfig() (m *MongoDB) {
	if _, err := toml.DecodeFile("C:\\Users\\ATG20\\go\\src\\GO-RMA\\core\\services\\mongo.toml", &m); err != nil {
		log.Fatal(err)
	}
	return
}

func MongoConnect() {
	var m = readMongoConfig()
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	Db = session.DB(m.Database)
}
