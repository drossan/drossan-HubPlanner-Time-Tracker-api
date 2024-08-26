package HubPlanner

type LoginResponse struct {
	Message      string `json:"message,omitempty"`
	Error        string `json:"error,omitempty"`
	Code         string `json:"code,omitempty"`
	Properties   string `json:"properties,omitempty"`
	Version      int    `json:"version,omitempty"`
	Status       bool   `json:"status,omitempty"`
	Location     string `json:"location,omitempty"`
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
