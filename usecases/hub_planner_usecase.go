package usecases

import (
	"hubplanner-proxy-api/domain/models/HubPlanner"
	"hubplanner-proxy-api/domain/repositories"
)

type HubPlannerUseCase struct {
	hubPlannerRepository repositories.HubPlannerRepository
}

func NewHubPlannerUserUseCase(hubPlannerRepository repositories.HubPlannerRepository) *HubPlannerUseCase {
	return &HubPlannerUseCase{
		hubPlannerRepository: hubPlannerRepository,
	}
}

func (uc *HubPlannerUseCase) Projects() ([]HubPlanner.Project, error) {
	return uc.hubPlannerRepository.Projects()
}

func (uc *HubPlannerUseCase) Categories() ([]HubPlanner.Category, error) {
	return uc.hubPlannerRepository.Categories()
}

func (uc *HubPlannerUseCase) TimeEntry(timeEntry *HubPlanner.TimeEntry) (*HubPlanner.TimeEntry, error) {
	return uc.hubPlannerRepository.TimeEntry(timeEntry)
}

func (uc *HubPlannerUseCase) TimeEntries(repositoryID string) (*HubPlanner.TimeEntries, error) {
	return uc.hubPlannerRepository.TimeEntries(repositoryID)
}

func (uc *HubPlannerUseCase) TimeEntrySubmit(timeEntryID, resourceID string) (*HubPlanner.TimeEntry, error) {
	return uc.hubPlannerRepository.TimeEntrySubmit(timeEntryID, resourceID)
}
