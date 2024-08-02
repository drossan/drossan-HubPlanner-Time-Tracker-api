package HubPlanner

import (
	"time"
)

type Resource struct {
	ID                            string        `json:"_id"`
	Email                         string        `json:"email"`
	Metadata                      string        `json:"metadata"`
	CreatedDate                   time.Time     `json:"createdDate"`
	UpdatedDate                   time.Time     `json:"updatedDate"`
	Note                          string        `json:"note"`
	FirstName                     string        `json:"firstName"`
	LastName                      string        `json:"lastName"`
	Status                        string        `json:"status"`
	Role                          string        `json:"role"`
	IsProjectManager              bool          `json:"isProjectManager"`
	Links                         Links         `json:"links"`
	Billing                       Billing       `json:"billing"`
	Billable                      bool          `json:"billable"`
	ResourceRates                 ResourceRates `json:"resourceRates"`
	Tags                          []string      `json:"tags"`
	IsApprover                    bool          `json:"isApprover"`
	UseCustomAvailability         bool          `json:"useCustomAvailability"`
	CustomVacationAllocationLimit float64       `json:"customVacationAllocationLimit"`
}

type Billing struct {
	UseDefault bool `json:"useDefault"`
	Rate       int  `json:"rate"`
}

type ResourceRates struct {
	External []interface{} `json:"external"`
	Internal []interface{} `json:"internal"`
}
