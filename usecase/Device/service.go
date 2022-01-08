package Device

import "github.com/AliceTrinta/GO-RMA/entity"

//Service  interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetDeviceByUUID(UUID string) (entity.Device, error) {
	return s.repo.GetByUUID(UUID)
}

func (s *Service) GetAllDevices() ([]entity.Device, error) {
	return s.repo.GetAll()
}

func (s *Service) CreateDevice(device entity.Device) error {
	return s.repo.Create(device)
}
