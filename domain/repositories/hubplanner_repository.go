package repositories

import (
	"hubplanner-proxy-api/domain/models/HubPlanner"
)

type HubPlannerRepository interface {
	Projects(resourceID string) ([]HubPlanner.Project, error)
	Categories() ([]HubPlanner.Category, error)
	TimeEntry(timeEntry *HubPlanner.TimeEntry) (*HubPlanner.TimeEntry, error)
	TimeEntrySubmit(timeEntryID, resourceID string) (*HubPlanner.TimeEntry, error)
	TimeEntries(repositoryID string) (*HubPlanner.TimeEntries, error)
}
