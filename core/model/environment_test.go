package model

import (
	"log"
	"testing"
)

//Testing the CreateEnvironment func.
func TestFakeCreateEnvironment(t *testing.T){
	var fake FakeEnvironment
	con := FakeConnection()
	err := fake.StartEnvironment(con)
	if err != nil {
		log.Fatal(err)
	}
}
