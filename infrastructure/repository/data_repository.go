package repository

import (
	"log"

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
	log.Println(data.DeviceUUID)
	err := m.db.C("data").Insert(data)
	return err
}

//GetAll function will return all data stored in the database.
func (m *DataMongo) GetAll() ([]entity.Data, error) {
	var data []entity.Data
	err := m.db.C("data").Find(nil).All(&data)
	return data, err
}
