package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing the StartDevice func.
func TestFakeDevice_StartDevice(t *testing.T) {
	var fake FakeDevice
	con := FakeConnection()
	err := fake.StartDevice(con)
	assert.Nil(t, err)
}
