package repository

import (
	"github.com/AliceTrinta/GO-RMA/entity"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
)

type EnvironmentMongo struct {
	db *mgo.Database
}

func NewEnvironmentMongo(db *mgo.Database) *EnvironmentMongo {
	return &EnvironmentMongo{
		db: db,
	}
}

//Create function will register an Environment in the database.
func (m *EnvironmentMongo) Create(env entity.Environment) error {
	err := m.db.C("environment").Insert(&env)
	return err
}

//GetByUUID function will search an Environment UUID in the database.
func (m *EnvironmentMongo) GetByUUID(UUID string) (entity.Environment, error) {
	var env entity.Environment
	err := m.db.C("environment").Find(bson.M{"UUID": UUID}).One(&env)
	return env, err
}
