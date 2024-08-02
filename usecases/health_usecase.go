package usecases

import (
	"hubplanner-proxy-api/domain/repositories"
)

type HealthUseCase struct {
	healthRepository repositories.HealthRepository
}

func NewHealthUseCase(healthRepository repositories.HealthRepository) *HealthUseCase {
	return &HealthUseCase{
		healthRepository: healthRepository,
	}
}

func (uc *HealthUseCase) CheckStatus() (map[string]string, error) {
	status, err := uc.healthRepository.Check()
	if err != nil {
		return status, err
	}
	return status, nil
}
