package manager

import (
	"lopei-grpc-server/service"

	//"github.com/go-delve/delve/service"
)

type ServiceManager interface {
	LopeiService() *service.LopeiService
}

type serviceManager struct {
	lopeiService *service.LopeiService
}

func (s *serviceManager) LopeiService() *service.LopeiService {
	return s.lopeiService
}

func NewServiceManager(repoManager RepoManager) ServiceManager {
	return &serviceManager{
		lopeiService: service.NewLopeiService(repoManager.LopeiRepo()),
	}
}
