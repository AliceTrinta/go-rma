package rml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing the forDevice func.
func TestForDevice(t *testing.T) {
	model.FakeConnection()
	model.DeviceInstance = model.FakeDevice{}
	err := forDevice("{\"UUID\":\"\",\"gatewayUUID\":\"\",\"name\":\"\",\"description\":\"\",\"cycleDelayInMillis\":\"\",\"resourceList\":[],\"status\":false}")
	assert.Nil(t, err)
}

//Testing the forData func.
func TestForData(t *testing.T) {
	model.FakeConnection()
	model.DataInstance = model.FakeData{}
	err := forData("{\"_id\":1,\"instant\":\"0001-01-01T00:00:00.000Z\",\"UUID\":\"\",\"resourceName\":\"\",\"value\":\"\"}")
	assert.Nil(t, err)
}

//Testing the forAction func.
func TestForAction(t *testing.T) {
	model.FakeConnection()
	model.ActionInstance = model.FakeAction{}
	err := forAction("{\"_id\":1,\"date\":\"0001-01-01T00:00:00.000Z\",\"UUID\":\"\",\"resourceName\":\"\",\"command\":\"\"}")
	assert.Nil(t, err)
}
