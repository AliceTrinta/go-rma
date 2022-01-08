package Device

import "github.com/AliceTrinta/GO-RMA/entity"

//Reader interface
type Reader interface {
	GetByUUID(UUID string) (entity.Device, error)
	GetAll() ([]entity.Device, error)
}

//Writer interface
type Writer interface {
	Create(device entity.Device) error
}

//Device Repository interface
type Repository interface {
	Reader
	Writer
}

//Device UseCase interface
type UseCase interface {
	GetDeviceByUUID(UUID string) (entity.Device, error)
	GetAllDevices() ([]entity.Device, error)
	CreateDevice(device entity.Device) error
}
