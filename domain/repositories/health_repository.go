package repositories

type HealthRepository interface {
	Check() (map[string]string, error)
}
