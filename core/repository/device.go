package repository

import (
	"GO-RMA/core/model"
)

func (m *MongoDB) GetDeviceByID(id string) (model.Device, error) {
	var device model.Device
	err := db.C("device").FindId(id).One(&device)
	return device, err
}

func (m *MongoDB) CreateDevice(device model.Device) error {
	err := db.C("device").Insert(&device)
	return err
}