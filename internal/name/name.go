package name

type Service struct {
	names []string
}

func NewService() *Service {
	return &Service{
		names: make([]string, 0),
	}
}

func (svc *Service) Add(name string) {
	svc.names = append(svc.names, name)
}

func (svc *Service) GetNames() []string {
	return svc.names
}
