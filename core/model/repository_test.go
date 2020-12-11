package model

import (
	"log"
	"testing"
)

//Testing the GetEnvironmentByUUID func.
func TestMongoDB_GetEnvironmentByUUID(t *testing.T) {
	con := FakeConnection()
	var UUID = ""
	_, err := con.GetEnvironmentByUUID(UUID)
	if err != nil{
		log.Fatal(err)
	}
}

//Testing the CreateEnvironment func.
func TestMongoDB_CreateEnvironment(t *testing.T) {
	var e Environment
	con := FakeConnection()
	err := con.CreateEnvironment(e)
	if err != nil {
		log.Fatal(err)
	}
}

//Testing the GetDeviceByUUID func.
func TestMongoDB_GetDeviceByUUID(t *testing.T) {
	con := FakeConnection()
	var UUID = ""
	_, err := con.GetDeviceByUUID(UUID)
	if err != nil {
		log.Fatal(err)
	}
}

//Testing the CreateDevice func.
func TestMongoDB_CreateDevice(t *testing.T) {
	var o Device
	con := FakeConnection()
	err := con.CreateDevice(o)
	if err != nil {
		log.Fatal(err)
	}
}

//Testing the CreateData func.
func TestMongoDB_CreateData(t *testing.T) {
	var d Data
	con := FakeConnection()
	err := con.CreateData(d)
	if err != nil {
		log.Fatal(err)
	}
}

//Testing the CreateAction func.
func TestMongoDB_CreateAction(t *testing.T) {
	var a Action
	con := FakeConnection()
	err := con.CreateAction(a)
	if err != nil {
		log.Fatal(err)
	}
}

