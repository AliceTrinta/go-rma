package Data

import "github.com/AliceTrinta/GO-RMA/entity"

//Reader interface
type Reader interface {
	GetAll() ([]entity.Data, error)
}

//Writer interface
type Writer interface {
	Create(data entity.Data) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	CreateData(data entity.Data) error
	GetAllData() ([]entity.Data, error)
}
