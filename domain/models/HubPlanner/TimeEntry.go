package HubPlanner

type TimeEntry struct {
	ID                 string `json:"_id,omitempty"`
	Resource           string `json:"resource,omitempty"`
	Project            string `json:"project"`
	ProjectName        string `json:"projectName,omitempty"`
	ProjectType        string `json:"projectType,omitempty"`
	ProjectStatus      string `json:"projectStatus,omitempty"`
	Date               string `json:"date"`
	Minutes            int    `json:"minutes"`
	Note               string `json:"note,omitempty"`
	CreatedDate        string `json:"createdDate,omitempty"`
	UpdatedDate        string `json:"updatedDate,omitempty"`
	Metadata           string `json:"metadata,omitempty"`
	Status             string `json:"status,omitempty"`
	Locked             bool   `json:"locked,omitempty"`
	Creator            string `json:"creator,omitempty"`
	CategoryTemplateId string `json:"categoryTemplateId,omitempty"`
	CategoryName       string `json:"categoryName,omitempty"`
	Billable           string `json:"billable,omitempty"`
}

type TimeEntryReduce struct {
	ID                 string `json:"_id,omitempty"`
	Project            string `json:"project"`
	ProjectName        string `json:"projectName,omitempty"`
	ProjectType        string `json:"projectType,omitempty"`
	Status             string `json:"status,omitempty"`
	Date               string `json:"date"`
	Minutes            int    `json:"minutes"`
	CategoryTemplateId string `json:"categoryTemplateId,omitempty"`
	CategoryName       string `json:"categoryName,omitempty"`
}

type TimeEntriesDayOfWeek struct {
	TotalTime string            `json:"total_time"`
	DayOfWeek string            `json:"day_of_week"`
	Items     []TimeEntryReduce `json:"items"`
}

type TimeEntriesWeek struct {
	TotalTime string                 `json:"total_time"`
	Week      string                 `json:"week"`
	Items     []TimeEntriesDayOfWeek `json:"items,omitempty"`
}

type TimeEntries struct {
	TotalItems int               `json:"totalItems"`
	Items      []TimeEntriesWeek `json:"items"`
}
