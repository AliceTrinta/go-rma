package model

import "go.mongodb.org/mongo-driver/bson"

//The GetEnvironmentByUUID function will search an Environment UUID in the database.
func (m *MongoDB) GetEnvironmentByUUID(UUID string) (Environment, error) {
	var env Environment
	err := Db.C("environment").Find(bson.M{"UUID": UUID}).One(&env)
	return env, err
}

//The CreateEnvironment function will register an Environment in the database.
func (m *MongoDB) CreateEnvironment(env Environment) error {
	err := Db.C("environment").Insert(&env)
	return err
}

//The GetAllDevices function will return all devices stored in the database.
func (m *MongoDB) GetAllDevices() ([]Device, error) {
	var devices []Device
	err := Db.C("device").Find(nil).All(&devices)
	return devices, err
}

//The GetDeviceByUUID function will search a Device UUID in the database.
func (m *MongoDB) GetDeviceByUUID(UUID string) (Device, error) {
	var device Device
	err := Db.C("device").Find(bson.M{"UUID": UUID}).One(&device)
	return device, err
}

//The CreateDevice function will register a Device in the database.
func (m *MongoDB) CreateDevice(device Device) error {
	err := Db.C("device").Insert(&device)
	return err
}

//The CreateData function will register a Data in the database.
func (m *MongoDB) CreateData(data Data) error {
	err := Db.C("data").Insert(data)
	return err
}

//The CreateAction function will register an Action in the database.
func (m *MongoDB) CreateAction(action Action) error {
	err := Db.C("action").Insert(action)
	return err
}
