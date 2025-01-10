package cli_models

import (
	"fmt"
)

// Структура для хранения контекста меню
type MenuContext struct {
	Breadcrumbs []string   // Хлебные крошки (путь)
	MenuItems   []MenuItem // Пункты меню
	isRunning   bool
}

// Метод для запуска контекста
func (c *MenuContext) Start() {
	c.isRunning = true
}

// Метод для проверки контекста
func (c *MenuContext) IsActive() bool {
	return c.isRunning
}

// Метод для остановки контекста
func (c *MenuContext) Stop() {
	c.isRunning = false
}

// Исполнение действия по id
func (c *MenuContext) RunAction(id int) {
	for _, item := range c.MenuItems {
		if item.ID == id {
			item.Action(c)
			return
		}
	}
	fmt.Println("Неверный выбор. Пожалуйста, выберите опцию из списка.")
}
