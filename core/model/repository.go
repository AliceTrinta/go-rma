package model

import "go.mongodb.org/mongo-driver/bson"

//GetEnvironmentByUUID function will search an Environment UUID in the database.
func (m *MongoDB) GetEnvironmentByUUID(UUID string) (Environment, error) {
	var env Environment
	err := Db.C("environment").Find(bson.M{"UUID": UUID}).One(&env)
	return env, err
}

//CreateEnvironment function will register an Environment in the database.
func (m *MongoDB) CreateEnvironment(env Environment) error {
	err := Db.C("environment").Insert(&env)
	return err
}

//GetAllDevices function will return all devices stored in the database.
func (m *MongoDB) GetAllDevices() ([]Device, error) {
	var devices []Device
	err := Db.C("device").Find(nil).All(&devices)
	return devices, err
}

//GetDeviceByUUID function will search a Device UUID in the database.
func (m *MongoDB) GetDeviceByUUID(UUID string) (Device, error) {
	var device Device
	err := Db.C("device").Find(bson.M{"UUID": UUID}).One(&device)
	return device, err
}

//CreateDevice function will register a Device in the database.
func (m *MongoDB) CreateDevice(device Device) error {
	err := Db.C("device").Insert(&device)
	return err
}

//CreateData function will register a Data in the database.
func (m *MongoDB) CreateData(data Data) error {
	err := Db.C("data").Insert(data)
	return err
}

//CreateAction function will register an Action in the database.
func (m *MongoDB) CreateAction(action Action) error {
	err := Db.C("action").Insert(action)
	return err
}
