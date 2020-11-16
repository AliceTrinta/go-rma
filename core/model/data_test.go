package model

import (
	"log"
	"testing"
)

func TestData_SaveData(t *testing.T) {
	var fake FakeData
	con := FakeConnection()
	err := fake.SaveData(con)
	if err != nil {
		log.Fatal(err)
	}
}
