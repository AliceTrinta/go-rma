package repository

import (
	"github.com/AliceTrinta/GO-RMA/entity"
	"go.mongodb.org/mongo-driver/bson"
)

//GetEnvironmentByUUID function will search an Environment UUID in the database.
func (m *MongoDB) GetEnvironmentByUUID(UUID string) (entity.Environment, error) {
	var env entity.Environment
	err := Db.C("environment").Find(bson.M{"UUID": UUID}).One(&env)
	return env, err
}

//CreateEnvironment function will register an Environment in the database.
func (m *MongoDB) CreateEnvironment(env entity.Environment) error {
	err := Db.C("environment").Insert(&env)
	return err
}

//GetAllDevices function will return all devices stored in the database.
func (m *MongoDB) GetAllDevices() ([]entity.Device, error) {
	var devices []entity.Device
	err := Db.C("device").Find(nil).All(&devices)
	return devices, err
}

//GetDeviceByUUID function will search a Device UUID in the database.
func (m *MongoDB) GetDeviceByUUID(UUID string) (entity.Device, error) {
	var device entity.Device
	err := Db.C("device").Find(bson.M{"UUID": UUID}).One(&device)
	return device, err
}

//CreateDevice function will register a Device in the database.
func (m *MongoDB) CreateDevice(device entity.Device) error {
	err := Db.C("device").Insert(&device)
	return err
}

//CreateData function will register a Data in the database.
func (m *MongoDB) CreateData(data entity.Data) error {
	err := Db.C("data").Insert(data)
	return err
}

//CreateAction function will register an Action in the database.
func (m *MongoDB) CreateAction(action entity.Action) error {
	err := Db.C("action").Insert(action)
	return err
}
