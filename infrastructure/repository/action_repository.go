package repository

import (
	"github.com/AliceTrinta/GO-RMA/entity"
	"gopkg.in/mgo.v2"
)

type ActionMongo struct {
	db *mgo.Database
}

func NewActionMongo(db *mgo.Database) *ActionMongo {
	return &ActionMongo{
		db: db,
	}
}

//Create function will register an Action in the database.
func (m *ActionMongo) Create(action entity.Action) error {
	err := m.db.C("action").Insert(action)
	return err
}
