package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing the CreateEnvironment func.
func TestFakeCreateEnvironment(t *testing.T) {
	var fake FakeEnvironment
	con := FakeConnection()
	err := fake.StartEnvironment(con)
	assert.Nil(t, err)
}
