package rml

import (
	"testing"

	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/stretchr/testify/assert"
)

//Testing the forDevice func.
func TestForDevice(t *testing.T) {
	entity.FakeConnection()
	entity.DeviceInstance = entity.FakeDevice{}
	err := forDevice("{\"UUID\":\"\",\"gatewayUUID\":\"\",\"name\":\"\",\"description\":\"\",\"cycleDelayInMillis\":\"\",\"resourceList\":[],\"status\":false}")
	assert.Nil(t, err)
}

//Testing the forData func.
func TestForData(t *testing.T) {
	entity.FakeConnection()
	entity.DataInstance = entity.FakeData{}
	err := forData("{\"_id\":1,\"instant\":\"0001-01-01T00:00:00.000Z\",\"UUID\":\"\",\"resourceName\":\"\",\"value\":\"\"}")
	assert.Nil(t, err)
}

//Testing the forAction func.
func TestForAction(t *testing.T) {
	entity.FakeConnection()
	entity.ActionInstance = entity.FakeAction{}
	err := forAction("{\"_id\":1,\"date\":\"0001-01-01T00:00:00.000Z\",\"UUID\":\"\",\"resourceName\":\"\",\"command\":\"\"}")
	assert.Nil(t, err)
}
