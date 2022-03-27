package main

type Service interface {
	Create() (string, error)
	GetLastCode() (int, error)
}

type serviceImp struct {
	repository repository
}

func NewService(repo repository) Service {
	return serviceImp{
		repository: repo,
	}
}

func (s serviceImp) Create() (string, error) {
	c := Code{}
	code := c.generate()
	err := s.repository.Create(code)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (s serviceImp) GetLastCode() (int, error) {
	code, err := s.repository.GetLastCode()
	if err != nil {
		return code, err
	}
	return code, nil
}
