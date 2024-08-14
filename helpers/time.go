package helpers

import (
	"fmt"

	"hubplanner-proxy-api/domain/models/HubPlanner"
)

// formatTotalTime formatea los minutos totales en una cadena con el formato adecuado.
func formatTotalTime(totalMinutes int) string {
	hours := totalMinutes / 60
	minutes := totalMinutes % 60

	if hours > 0 && minutes > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	} else if hours > 0 {
		return fmt.Sprintf("%dh", hours)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm", minutes)
	} else {
		return "0m"
	}
}

// CalculateTotalTime calcula el tiempo total de una lista de entradas de tiempo.
func CalculateTotalTime(entries []HubPlanner.TimeEntryReduce) (string, int) {
	totalMinutes := 0
	for _, entry := range entries {
		totalMinutes += entry.Minutes
	}
	return formatTotalTime(totalMinutes), totalMinutes
}

// CalculateTotalTimeForWeek calcula el tiempo total de una semana.
func CalculateTotalTimeForWeek(entries []HubPlanner.TimeEntriesDayOfWeek) (string, int) {
	totalMinutes := 0
	for _, entry := range entries {
		for _, item := range entry.Items {
			totalMinutes += item.Minutes
		}
	}
	return formatTotalTime(totalMinutes), totalMinutes
}
