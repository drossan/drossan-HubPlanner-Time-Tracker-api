package repository

import "hubplanner-proxy-api/domain/repositories"

type healthConnectionRepository struct {
}

func NewHealthConnectionRepository() repositories.HealthRepository {
	return &healthConnectionRepository{}
}

func (r *healthConnectionRepository) Check() (map[string]string, error) {
	return map[string]string{"status": "ok"}, nil
}
