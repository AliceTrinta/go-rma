package repository

import (
	"log"

	"github.com/AliceTrinta/GO-RMA/entity"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
)

type DeviceMongo struct {
	db *mgo.Database
}

func NewDeviceMongo(db *mgo.Database) *DeviceMongo {
	return &DeviceMongo{
		db: db,
	}
}

//Create function will register a Device in the database.
func (m *DeviceMongo) Create(device entity.Device) error {
	err := m.db.C("device").Insert(&device)
	return err
}

//GetByUUID function will search a Device UUID in the database.
func (m *DeviceMongo) GetByUUID(UUID string) (entity.Device, error) {
	var device entity.Device
	err := m.db.C("device").Find(bson.M{"UUID": UUID}).One(&device)
	log.Println(err.Error())
	return device, err
}

//GetAll function will return all devices stored in the database.
func (m *DeviceMongo) GetAll() ([]entity.Device, error) {
	var devices []entity.Device
	err := m.db.C("device").Find(nil).All(&devices)
	return devices, err
}
