package repository

import "GO-RMA/core/model"

func (m *MongoDB) CreateAction(action model.Action) error {
	err := db.C("action").Insert(action)
	return err
}
