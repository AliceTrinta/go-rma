package model

import (
	"log"
	"testing"
)

func TestMongoDB_GetDeviceByID(t *testing.T) {
	con := FakeConnection()
	var id = "device"
	_, err := con.GetDeviceByID(id)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMongoDB_CreateDevice(t *testing.T) {
	var o Device
	con := FakeConnection()
	err := con.CreateDevice(o)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMongoDB_CreateData(t *testing.T) {
	var d Data
	con := FakeConnection()
	err := con.CreateData(d)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMongoDB_CreateAction(t *testing.T) {
	var a Action
	con := FakeConnection()
	err := con.CreateAction(a)
	if err != nil {
		log.Fatal(err)
	}
}

