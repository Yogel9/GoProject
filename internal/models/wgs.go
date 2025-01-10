package models

// WasteGenerationSite представляет место образования отходов
type WasteGenerationSite struct {
	ID      int
	Name    string
	Address string
	Wastes  []Waste // Список отходов на месте образования
}
