package repository

import (
	"testing"

	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/stretchr/testify/assert"
)

func FakeEnvironmentMongo() (con EnvironmentMongo) {
	con.db = MongoConnect()
	return con
}

//Testing the GetByUUID func from Environment.
func TestEnvironmentRepository_GetByUUID(t *testing.T) {
	con := FakeEnvironmentMongo()
	var UUID = ""
	_, err := con.GetByUUID(UUID)
	assert.Nil(t, err)
}

//Testing the Create func from Environment.
func TestEnvironmentRepository_Create(t *testing.T) {
	var e entity.Environment
	con := FakeEnvironmentMongo()
	err := con.Create(e)
	assert.Nil(t, err)
}
