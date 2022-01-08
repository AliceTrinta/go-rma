package repository

import (
	"testing"

	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/stretchr/testify/assert"
)

//Testing the GetEnvironmentByUUID func.
func TestMongoDB_GetEnvironmentByUUID(t *testing.T) {
	con := FakeConnection()
	var UUID = ""
	_, err := con.GetEnvironmentByUUID(UUID)
	assert.Nil(t, err)
}

//Testing the CreateEnvironment func.
func TestMongoDB_CreateEnvironment(t *testing.T) {
	var e entity.Environment
	con := FakeConnection()
	err := con.CreateEnvironment(e)
	assert.Nil(t, err)
}

//Testing the GetAllDevices func.
func TestMongoDB_GetAllDevices(t *testing.T) {
	con := FakeConnection()
	_, err := con.GetAllDevices()
	assert.Nil(t, err)
}

//Testing the GetDeviceByUUID func.
func TestMongoDB_GetDeviceByUUID(t *testing.T) {
	con := FakeConnection()
	var UUID = ""
	_, err := con.GetDeviceByUUID(UUID)
	assert.Nil(t, err)
}

//Testing the CreateDevice func.
func TestMongoDB_CreateDevice(t *testing.T) {
	var o entity.Device
	con := FakeConnection()
	err := con.CreateDevice(o)
	assert.Nil(t, err)
}

//Testing the CreateData func.
func TestMongoDB_CreateData(t *testing.T) {
	var d entity.Data
	con := FakeConnection()
	err := con.CreateData(d)
	assert.Nil(t, err)
}

//Testing the CreateAction func.
func TestMongoDB_CreateAction(t *testing.T) {
	var a entity.Action
	con := FakeConnection()
	err := con.CreateAction(a)
	assert.Nil(t, err)
}
