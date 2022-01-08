package repository

import (
	"testing"

	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/stretchr/testify/assert"
)

func FakeDeviceMongo() (con DeviceMongo) {
	con.db = MongoConnect()
	return con
}

//Testing the GetAll func from Device.
func TestDeviceRepository_GetAll(t *testing.T) {
	con := FakeDeviceMongo()
	_, err := con.GetAll()
	assert.Nil(t, err)
}

//Testing the GetByUUID func from Device.
func TestDeviceRepository_GetByUUID(t *testing.T) {
	con := FakeDeviceMongo()
	var UUID = ""
	_, err := con.GetByUUID(UUID)
	assert.Nil(t, err)
}

//Testing the Create func from Device.
func TestDeviceRepository_Create(t *testing.T) {
	var o entity.Device
	con := FakeDeviceMongo()
	err := con.Create(o)
	assert.Nil(t, err)
}
