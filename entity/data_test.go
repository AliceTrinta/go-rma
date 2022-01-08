package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing the SaveData func.
func TestData_SaveData(t *testing.T) {
	var fake FakeData
	con := FakeConnection()
	err := fake.SaveData(con)
	assert.Nil(t, err)
}
