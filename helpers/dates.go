package helpers

import (
	"time"
)

type WeekRange struct {
	Week  string   `json:"week"`
	Dates []string `json:"dates"`
}

func GetWeekRanges() []WeekRange {
	var weekRanges []WeekRange
	dateFormat := "2006-01-02"

	currentDate := time.Now()
	weekday := int(currentDate.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	startOfWeek := currentDate.AddDate(0, 0, -weekday+1) // Adjust to start from Monday

	for i := 0; i < 4; i++ {
		weekStart := startOfWeek.AddDate(0, 0, -7*i)
		weekEnd := weekStart.AddDate(0, 0, 6)

		var weekLabel string
		if i == 0 {
			weekLabel = "Esta semana"
		} else {
			weekLabel = weekStart.Format("02 Jan") + " - " + weekEnd.Format("02 Jan")
		}

		var dates []string
		for j := 0; j < 7; j++ {
			date := weekStart.AddDate(0, 0, j)
			if date.After(currentDate) {
				break
			}
			dates = append(dates, date.Format(dateFormat))
		}
		weekRanges = append(weekRanges, WeekRange{Week: weekLabel, Dates: dates})
	}

	return weekRanges
}
