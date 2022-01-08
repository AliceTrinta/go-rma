package repository

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

//Testing the MongoConnect func.
func TestMongoConnect(t *testing.T) {
	MongoConnect()
}

//Testing the ReadMongoConfig func.
func TestReadMongoConfig(t *testing.T) {
	m, err := ReadMongoConfig()
	assert.Nil(t, err)
	log.Println(m)
}

//Testing the Db variable.
func TestDbVariable(t *testing.T) {
	m, err := ReadMongoConfig()
	assert.Nil(t, err)
	session, err := mgo.Dial(m.Server)
	assert.Nil(t, err)
	Db = session.DB(m.Database)
	if Db == nil {
		assert.Nil(t, Db)
	}
}
