package repositories

// BaseRepository определяет методы для работы
type BaseRepository interface {
	SaveWaste()    // Сохранить отходы
	GetWasteByID() // Получить отходы по ID
	GetAllWastes() // Получить все отходы
	DeleteWaste()  // Удалить отходы
}
