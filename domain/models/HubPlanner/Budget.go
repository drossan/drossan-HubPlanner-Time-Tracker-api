package HubPlanner

type Budget struct {
	HasBudget    bool         `json:"hasBudget"`
	ProjectHours ProjectHours `json:"projectHours"`
	CashAmount   CashAmount   `json:"cashAmount"`
}

type BudgetCategory struct {
	BudgetHours      int    `json:"budgetHours"`
	BudgetCashAmount int    `json:"budgetCashAmount"`
	BudgetCurrency   string `json:"budgetCurrency"`
	CategoryID       string `json:"categoryId"`
	ID               string `json:"_id"`
}
