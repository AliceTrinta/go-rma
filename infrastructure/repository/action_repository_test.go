package repository

import (
	"testing"

	"github.com/AliceTrinta/GO-RMA/config"
	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/stretchr/testify/assert"
)

func FakeActionMongo() (con ActionMongo) {
	db, err := config.MongoConnect()
	con.db = db
	if err != nil {
		return
	}
	return con
}

//Testing the Create func from Action.
func TestActionRepository_Create(t *testing.T) {
	var a entity.Action
	con := FakeActionMongo()
	err := con.Create(a)
	assert.Nil(t, err)
}
