package repository

import (
	"github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2"
	"log"
)

type MongoDB struct {
	Server string
	Database string
}

var db *mgo.Database

func readMongoConfig() (m *MongoDB) {
	if _, err := toml.DecodeFile("core/repository/mongo.toml", &m); err != nil {
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
	db = session.DB(m.Database)
}
