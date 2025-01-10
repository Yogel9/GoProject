package models

// Waste представляет собой отходы
type Waste struct {
	ID          int
	Name        string  // Название отхода
	Quantity    float64 // Количество отходов
	HazardClass int     // Класс опасности
}
