package repository

import (
	"testing"

	"github.com/AliceTrinta/GO-RMA/config"
	"github.com/AliceTrinta/GO-RMA/entity"
	"github.com/stretchr/testify/assert"
)

func FakeDataMongo() (con DataMongo) {
	db, err := config.MongoConnect()
	con.db = db
	if err != nil {
		return
	}
	return con
}

//Testing the Create func from Data.
func TestMongoDB_CreateData(t *testing.T) {
	var d entity.Data
	con := FakeDataMongo()
	err := con.Create(d)
	assert.Nil(t, err)
}
