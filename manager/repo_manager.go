package manager

import "lopei-grpc-server/repo"

type RepoManager interface {
	LopeiRepo() repo.LopeiRepo
}

type repoManager struct {
	lopeiRepo repo.LopeiRepo
}

func (r *repoManager) LopeiRepo() repo.LopeiRepo {
	return r.lopeiRepo
}

func NewRepoManager() RepoManager {
	reposs := new(repoManager)
	reposs.lopeiRepo = repo.NewLopeiRepo()
	return reposs
}
