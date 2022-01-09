package config

import (
	"log"

	"gopkg.in/mgo.v2"
)

//MongoConnect function will establish the connection with the database.
func MongoConnect() (*mgo.Database, error) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return session.DB(DATABASE), nil
}
