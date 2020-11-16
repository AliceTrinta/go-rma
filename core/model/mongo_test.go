package model

import (
	"log"
	"testing"
)

func TestMongoConnect(t *testing.T) {
	MongoConnect()
}

func TestReadMongoConfig(t *testing.T) {
	m := ReadMongoConfig()
	log.Println(m)
}