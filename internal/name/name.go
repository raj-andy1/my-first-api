package name

import "errors"

type Service struct {
	names []string
}

func NewService() *Service {
	return &Service{
		names: make([]string, 0),
	}
}

func (svc *Service) Add(name string) error {
	for _, n := range svc.names {
		if n == name {
			return errors.New("service name is already in use")
		} else {
			svc.names = append(svc.names, name)
		}
	}
	return nil
}

func (svc *Service) GetNames() []string {
	return svc.names
}
