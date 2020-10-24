package repository

import "GO-RMA/core/model"

func (m *MongoDB) CreateData(data model.Data) error {
	err := db.C("data").Insert(data)
	return err
}
