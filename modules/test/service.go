package test

type Service interface {
	GetTest(input string) string
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTest(input string) string {
	test := s.repository.GetTest(input)
	return test
}
