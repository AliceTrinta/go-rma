package repository

import (
	"github.com/AliceTrinta/GO-RMA/entity"
	"gopkg.in/mgo.v2"
)

type DataMongo struct {
	db *mgo.Database
}

func NewDataMongo(db *mgo.Database) *DataMongo {
	return &DataMongo{
		db: db,
	}
}

//Create function will register a Data in the database.
func (m *DataMongo) Create(data entity.Data) error {
	err := m.db.C("data").Insert(data)
	return err
}
