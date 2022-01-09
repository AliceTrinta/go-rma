package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing the MongoConnect func.
func TestMongoConnect(t *testing.T) {
	_, err := MongoConnect()
	assert.Nil(t, err)
}
