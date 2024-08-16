package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"hubplanner-proxy-api/domain/models/HubPlanner"
	"hubplanner-proxy-api/domain/repositories"
	"hubplanner-proxy-api/helpers"
)

type hubPlannerAPIRepository struct{}

func NewHubPlannerConnectionRepository() repositories.HubPlannerRepository {
	return &hubPlannerAPIRepository{}
}

func (r *hubPlannerAPIRepository) Projects(resourceID string) ([]HubPlanner.Project, error) {
	var projects []HubPlanner.Project

	url := os.Getenv("API_URL") + "/project/search?sort=-createdDate&limit=1000"
	method := "POST"

	payload := fmt.Sprintf(`{
		"resources": {
			"$in": ["%s"]
		}
	}`, resourceID)

	bodyBytes, err := helpers.MakeHTTPRequest(
		method,
		url,
		os.Getenv("API_TOKEN"),
		"application/json",
		strings.NewReader(payload),
	)

	if err != nil {
		return projects, err
	}

	_ = json.Unmarshal(bodyBytes, &projects)

	events, err := r.events()
	if err != nil {
		return projects, err
	}

	projects = append(projects, events...)

	return projects, nil
}

func (r *hubPlannerAPIRepository) events() ([]HubPlanner.Project, error) {
	var events []HubPlanner.Project

	url := os.Getenv("API_URL") + "/event/search?sort=-createdDate&limit=1000"
	method := "POST"

	bodyBytes, err := helpers.MakeHTTPRequest(
		method,
		url,
		os.Getenv("API_TOKEN"),
		"application/json",
		nil,
	)

	if err != nil {
		return events, err
	}

	_ = json.Unmarshal(bodyBytes, &events)
	return events, nil
}

func (r *hubPlannerAPIRepository) Categories() ([]HubPlanner.Category, error) {
	var categories []HubPlanner.Category

	url := os.Getenv("API_URL") + "/categories"
	method := "GET"

	bodyBytes, err := helpers.MakeHTTPRequest(
		method,
		url,
		os.Getenv("API_TOKEN"),
		"application/json",
		nil,
	)
	if err != nil {
		return categories, err
	}

	_ = json.Unmarshal(bodyBytes, &categories)
	return categories, nil
}

func (r *hubPlannerAPIRepository) TimeEntry(timeEntry *HubPlanner.TimeEntry) (*HubPlanner.TimeEntry, error) {
	// NOTE: Recuperar entradas por la fecha recibida
	timeEntries, err := r.recoveryTimeEntriesByDate(timeEntry.Resource, timeEntry.Date)

	// NOTE: Recorrer timeEntries y buscar si existe el proyecto que y la categoría que hemos recibido en timeEntry
	for _, entry := range timeEntries {
		if entry.Project == timeEntry.Project && entry.CategoryTemplateId == timeEntry.CategoryTemplateId {
			timeEntry.ID = entry.ID
			timeEntry.ProjectName = entry.ProjectName
			timeEntry.ProjectType = entry.ProjectType
			timeEntry.ProjectStatus = entry.ProjectStatus
			timeEntry.Minutes = timeEntry.Minutes + entry.Minutes
			timeEntry.CreatedDate = entry.CreatedDate
			timeEntry.UpdatedDate = entry.UpdatedDate
			timeEntry.Metadata = entry.Metadata
			timeEntry.Status = entry.Status
			timeEntry.Locked = entry.Locked
			timeEntry.Creator = entry.Creator
			timeEntry.CategoryName = entry.CategoryName
			timeEntry.Billable = entry.Billable

			if entry.Note != "" {
				timeEntry.Note = entry.Note + "\n" + timeEntry.Note
			}
		}
	}

	timeEntry.Status = "UNSUBMITTED"
	timeEntry, err = r.addTimeEntry(timeEntry)
	if err != nil {
		return timeEntry, err
	}

	return timeEntry, nil
}

func (r *hubPlannerAPIRepository) recoveryTimeEntriesByDate(resourceID, date string) ([]HubPlanner.TimeEntry, error) {
	var timeEntries []HubPlanner.TimeEntry

	url := os.Getenv("API_URL") + "/timeentry/search"
	method := "POST"

	data := map[string]string{
		"resource": resourceID,
		"date":     date,
	}
	payload, _ := json.Marshal(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		return timeEntries, err
	}
	req.Header.Add("Authorization", os.Getenv("API_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Secuoyas Experiences - Time Tracking (daniel.rossello@secuoyas.com)")

	res, err := client.Do(req)
	if err != nil {
		return timeEntries, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return timeEntries, err
	}

	_ = json.Unmarshal(body, &timeEntries)
	return timeEntries, nil
}

func (r *hubPlannerAPIRepository) recoveryTimeEntriesByID(timeEntryID string) (*HubPlanner.TimeEntry, error) {
	var timeEntry HubPlanner.TimeEntry

	url := os.Getenv("API_URL") + "/timeentry/" + timeEntryID
	method := "GET"

	bodyBytes, err := helpers.MakeHTTPRequest(
		method,
		url,
		os.Getenv("API_TOKEN"),
		"application/json",
		nil,
	)
	if err != nil {
		return &timeEntry, err
	}

	_ = json.Unmarshal(bodyBytes, &timeEntry)
	return &timeEntry, nil
}

func (r *hubPlannerAPIRepository) addTimeEntry(timeEntry *HubPlanner.TimeEntry) (*HubPlanner.TimeEntry, error) {
	url := os.Getenv("API_URL") + "/timeentry"
	method := "POST"

	// NOTE: Si existe timeEntry.ID mandamos una actualización, de lo contrario es un POST
	if timeEntry.ID != "" {
		url = os.Getenv("API_URL") + "/timeentry/" + timeEntry.ID
		method = "PUT"
	}

	payloadBytes, err := json.Marshal(timeEntry)
	if err != nil {
		return timeEntry, err
	}
	payload := bytes.NewBuffer(payloadBytes)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return timeEntry, err
	}
	req.Header.Add("Authorization", os.Getenv("API_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Secuoyas Experiences - Time Tracking (daniel.rossello@secuoyas.com)")

	res, err := client.Do(req)
	if err != nil {
		return timeEntry, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return timeEntry, err
	}

	_ = json.Unmarshal(body, &timeEntry)
	return timeEntry, nil
}

func (r *hubPlannerAPIRepository) TimeEntries(resourceID string) (*HubPlanner.TimeEntries, error) {
	weekRanges := helpers.GetWeekRanges()
	fmt.Println(weekRanges)

	timeEntries := HubPlanner.TimeEntries{}

	for _, weekRange := range weekRanges {
		// Obtener la fecha de inicio y fin de la semana
		weekDateStart := weekRange.Dates[0]
		weekDateEnd := weekRange.Dates[len(weekRange.Dates)-1]

		timeEntriesWeek := HubPlanner.TimeEntriesWeek{
			Week: weekRange.Week,
		}

		// Crear un mapa para agrupar las entradas por día
		dailyEntries := make(map[string][]HubPlanner.TimeEntryReduce)

		page := 0
		for {
			entries, err := r.recoveryTimeEntriesByWeek(resourceID, weekDateStart, weekDateEnd, fmt.Sprintf("%d", page))
			if err != nil || len(entries) == 0 {
				break
			}

			// Filtrar y agrupar entradas por día
			for _, entry := range entries {
				if entry.Minutes > 0 {
					dailyEntries[entry.Date] = append(dailyEntries[entry.Date], entry)
				}
			}

			page++
		}

		// Crear las estructuras de los días de la semana con las entradas agrupadas
		for day, entries := range dailyEntries {
			timeEntriesDayOfWeek := HubPlanner.TimeEntriesDayOfWeek{
				DayOfWeek: day,
				Items:     entries,
			}

			// Agregar total de tiempo por día de la semana
			timeEntriesDayOfWeek.TotalTime, timeEntriesDayOfWeek.TotalTimeInMinutes = helpers.CalculateTotalTime(timeEntriesDayOfWeek.Items)
			timeEntriesWeek.Items = append(timeEntriesWeek.Items, timeEntriesDayOfWeek)
		}

		// Agregar total de tiempo por semana
		timeEntriesWeek.TotalTime, timeEntriesWeek.TotalTimeInMinutes = helpers.CalculateTotalTimeForWeek(timeEntriesWeek.Items)
		timeEntries.Items = append(timeEntries.Items, timeEntriesWeek)
	}

	timeEntries.TotalItems = len(timeEntries.Items)

	return &timeEntries, nil
}

func (r *hubPlannerAPIRepository) recoveryTimeEntriesByWeek(resourceID, weekDateStart, weekDateEnd, page string) ([]HubPlanner.TimeEntryReduce, error) {
	var timeEntries []HubPlanner.TimeEntryReduce

	url := os.Getenv("API_URL") + "/timeentry/search?page=" + page + "&limit=100"
	method := "POST"

	data := map[string]interface{}{
		"resource": resourceID,
		"date": map[string]interface{}{
			"$gte": weekDateStart,
			"$lte": weekDateEnd,
		},
	}
	payload, _ := json.Marshal(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		return timeEntries, err
	}
	req.Header.Add("Authorization", os.Getenv("API_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Secuoyas Experiences - Time Tracking (daniel.rossello@secuoyas.com)")

	res, err := client.Do(req)
	if err != nil {
		return timeEntries, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return timeEntries, err
	}

	_ = json.Unmarshal(body, &timeEntries)
	return timeEntries, nil
}

func (r *hubPlannerAPIRepository) TimeEntrySubmit(timeEntryID, resourceID string) (*HubPlanner.TimeEntry, error) {
	timeEntry, err := r.recoveryTimeEntriesByID(timeEntryID)

	if timeEntry.Resource != resourceID {
		if err != nil {
			return nil, errors.New("la entrada no corresponde al usuario")
		}
	}

	timeEntry.Status = "SUBMITTED"
	timeEntry, err = r.addTimeEntry(timeEntry)

	if err != nil {
		return nil, err
	}

	return timeEntry, nil
}
