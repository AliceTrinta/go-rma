package Data

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

func (s *Service) CreateData(data entity.Data) error {
	return s.repo.Create(data)
}

func (s *Service) GetAllData() ([]entity.Data, error) {
	return s.repo.GetAll()
}