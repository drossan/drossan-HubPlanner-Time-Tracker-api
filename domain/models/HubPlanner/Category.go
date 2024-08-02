package HubPlanner

import "time"

type Category struct {
	ID          string    `json:"_id"`
	Name        string    `json:"name"`
	GridColor   string    `json:"gridColor"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	Type        string    `json:"type"`
}
