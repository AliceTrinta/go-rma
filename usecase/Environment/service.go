package Environment

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

func (s *Service) GetEnvironmentByUUID(UUID string) (entity.Environment, error) {
	return s.repo.GetByUUID(UUID)
}

func (s *Service) CreateEnvironment(env entity.Environment) error {
	return s.repo.Create(env)
}
