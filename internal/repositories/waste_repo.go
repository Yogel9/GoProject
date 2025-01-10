package repositories

import "GoProject/internal/models"

// WasteRepository определяет методы для работы с отходами
type WasteRepository interface {
	SaveWaste(waste models.Waste) error         // Сохранить отходы
	GetWasteByID(id int) (*models.Waste, error) // Получить отходы по ID
	GetAllWastes() ([]models.Waste, error)      // Получить все отходы
	DeleteWaste(id int) error                   // Удалить отходы
}
