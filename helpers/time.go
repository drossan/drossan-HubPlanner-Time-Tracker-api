package helpers

import (
	"fmt"

	"hubplanner-proxy-api/domain/models/HubPlanner"
)

// CalculateTotalTime calcula el tiempo total de una lista de entradas de tiempo.
func CalculateTotalTime(entries []HubPlanner.TimeEntryReduce) string {
	totalMinutes := 0
	for _, entry := range entries {
		totalMinutes += entry.Minutes
	}
	return fmt.Sprintf("%dh %dm", totalMinutes/60, totalMinutes%60)
}

// CalculateTotalTimeForWeek calcula el tiempo total de una semana.
func CalculateTotalTimeForWeek(entries []HubPlanner.TimeEntriesDayOfWeek) string {
	totalMinutes := 0
	for _, entry := range entries {
		for _, item := range entry.Items {
			totalMinutes += item.Minutes
		}
	}
	return fmt.Sprintf("%dh %dm", totalMinutes/60, totalMinutes%60)
}
