package model

import (
	"log"
	"testing"
)

//Testing the SaveData func.
func TestData_SaveData(t *testing.T) {
	var fake FakeData
	con := FakeConnection()
	err := fake.SaveData(con)
	if err != nil {
		log.Fatal(err)
	}
}
