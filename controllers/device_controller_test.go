package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing the Create func from Action.
func TestDeviceController_StartDevice(t *testing.T) {
	_, err := StartDevice([]byte ("{\"UUID\":\"\",\"gatewayUUID\":\"\",\"name\":\"\",\"description\":\"\",\"cycleDelayInMillis\":\"\",\"resourceList\":[],\"status\":false}"))
	assert.Nil(t, err)
}
