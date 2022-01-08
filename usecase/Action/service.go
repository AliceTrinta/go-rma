package Action

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

func (s *Service) CreateAction(action entity.Action) error {
	return s.repo.Create(action)
}
