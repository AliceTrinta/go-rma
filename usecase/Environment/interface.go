package Environment

import "github.com/AliceTrinta/GO-RMA/entity"

//Reader interface
type Reader interface {
	GetByUUID(UUID string) (entity.Environment, error)
}

//Writer interface
type Writer interface {
	Create(env entity.Environment) error
}

//Device Repository interface
type Repository interface {
	Reader
	Writer
}

//Device UseCase interface
type UseCase interface {
	GetEnvironmentByUUID(UUID string) (entity.Environment, error)
	CreateEnvironment(env entity.Environment) error
}
