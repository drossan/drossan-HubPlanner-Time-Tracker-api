package HubPlanner

type Project struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	//Links                    Links            `json:"links"`
	//Note              string    `json:"note"`
	//CreatedDate       time.Time `json:"createdDate"`
	//UpdatedDate       time.Time `json:"updatedDate"`
	//TimeEntryEnabled  bool      `json:"timeEntryEnabled"`
	//TimeEntryLocked   bool      `json:"timeEntryLocked"`
	//TimeEntryApproval bool      `json:"timeEntryApproval"`
	//ProjectCode       string    `json:"projectCode"`
	//Tags              []string  `json:"tags"`
	//TimeEntryNoteRequired    bool             `json:"timeEntryNoteRequired"`
	//WorkDays []bool `json:"workDays"`
	//UseProjectDays           bool             `json:"useProjectDays"`
	//BudgetCategories         []BudgetCategory `json:"budgetCategories"`
	//FixedCosts               []interface{}    `json:"fixedCosts"`
	//Budget                   Budget           `json:"budget"`
	//BudgetHours              int              `json:"budgetHours"`
	//BudgetCashAmount         int              `json:"budgetCashAmount"`
	//BudgetCurrency           string           `json:"budgetCurrency"`
	//CompanyBillingRateID     interface{}      `json:"companyBillingRateId"`
	//ResourceRates            []interface{}    `json:"resourceRates"`
	//UseStatusColor           bool             `json:"useStatusColor"`
	//Status                   string           `json:"status"`
	//UseProjectDuration       bool             `json:"useProjectDuration"`
	//Start                    interface{}      `json:"start"`
	//End                      interface{}      `json:"end"`
	//IncludeBookedTimeReports bool             `json:"includeBookedTimeReports"`
	//IncludeBookedTimeGrid    bool             `json:"includeBookedTimeGrid"`
	//Private                  bool             `json:"private"`
	//ProjectManagers          []interface{}    `json:"projectManagers"`
	//Resources                []string         `json:"resources"`
	//BackgroundColor          string           `json:"backgroundColor"`
	//Metadata                 string           `json:"metadata"`
	//Billable                 bool             `json:"billable"`
	//ProjectRate              ProjectRate      `json:"projectRate"`
	//Customers                []Customer       `json:"customers"`
	//DefaultCategory          string           `json:"defaultCategory"`
}

type ProjectHours struct {
	Active bool `json:"active"`
	Hours  int  `json:"hours"`
}

type CashAmount struct {
	Active      bool        `json:"active"`
	Amount      int         `json:"amount"`
	Currency    string      `json:"currency"`
	BillingRate BillingRate `json:"billingRate"`
}

type BillingRate struct {
	UseDefault bool        `json:"useDefault"`
	Rate       int         `json:"rate"`
	ID         interface{} `json:"id"`
}

type ProjectRate struct {
	External Rate `json:"external"`
	Internal Rate `json:"internal"`
}

type Rate struct {
	DefaultRateID interface{}   `json:"defaultRateId"`
	CustomRates   []interface{} `json:"customRates"`
}

type Customer struct {
	CustomerID string `json:"customerId"`
	Name       string `json:"name"`
	ID         string `json:"_id"`
}
