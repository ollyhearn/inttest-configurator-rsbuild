package projects

import "go.uber.org/zap"

type UseCase struct {
	log      *zap.SugaredLogger
	projRepo ProjectRepository
}

func New(projectRepo ProjectRepository, log *zap.SugaredLogger) *UseCase {
	return &UseCase{
		log:      log,
		projRepo: projectRepo,
	}
}
