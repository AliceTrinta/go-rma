package model

import (
	"gopkg.in/mgo.v2"
	"log"
	"testing"
)

//Testing the MongoConnect func.
func TestMongoConnect(t *testing.T) {
	MongoConnect()
}

//Testing the ReadMongoConfig func.
func TestReadMongoConfig(t *testing.T) {
	m, err:= ReadMongoConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(m)
}

//Testing the Db variable.
func TestDbVariable(t *testing.T){
	m, err := ReadMongoConfig()
	if err != nil {
		log.Fatal(err)
	}
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	Db = session.DB(m.Database)
	if Db == nil {
		log.Fatal("Db variable is empty.")
	}
}