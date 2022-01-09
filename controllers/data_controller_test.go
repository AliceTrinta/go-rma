package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing the Create func from Action.
func TestDataController_SaveData(t *testing.T) {
	err := SaveData([]byte ("{\" id\":1,\"instant\":\"0001-01-01T00:00:00.000Z\",\"UUID\":\"\",\"resourceName\":\"\",\"value\":\"\"}"))
	assert.Nil(t, err)
}
