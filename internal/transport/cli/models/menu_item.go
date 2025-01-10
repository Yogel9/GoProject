package cli_models

type MenuItem struct {
	ID     int
	Label  string
	Action func(*MenuContext)
}
