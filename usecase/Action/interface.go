package Action

import "github.com/AliceTrinta/GO-RMA/entity"

//Reader interface
type Reader interface {
	//In progress...
}

//Writer user writer
type Writer interface {
	Create(action entity.Action) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	CreateAction(action entity.Action) error
}
