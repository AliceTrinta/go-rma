package model

import (
	"log"
	"testing"
)

//Testing the StartDevice func.
func TestFakeDevice_StartDevice(t *testing.T) {
	var fake FakeDevice
	con := FakeConnection()
	err := fake.StartDevice(con)
	if err != nil {
		log.Fatal(err)
	}
}