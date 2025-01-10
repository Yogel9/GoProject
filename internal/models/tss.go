package models

// TemporaryStorageSite представляет площадку временного хранения отходов
type TemporaryStorageSite struct {
	ID      int
	Name    string
	Address string
	Wastes  []Waste
}
