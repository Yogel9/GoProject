package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Метод для остановки контекста
func (c *MenuContext) Stop() {
	c.isRunning = false
}

type MenuItem struct {
	ID     int
	Label  string
	Action func(*MenuContext)
}

// Функция для отображения меню
func displayMenu(context *MenuContext) {
	fmt.Println("=== Контекстное меню ===")
	fmt.Printf("Path: %v\n", strings.Join(context.Breadcrumbs, " >> "))
	for _, item := range context.MenuItems {
		fmt.Printf("[%d]. %s\n", item.ID, item.Label)
	}
	fmt.Print("Выберите опцию: ")
}

// Функция для получения выбора пользователя
func getUserChoice() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		return -1 // Возвращаем -1, если ввод некорректен
	}
	return choice
}

func getMenuContext(context *MenuContext) {
	if len(context.Breadcrumbs) == 0 {
		context.Breadcrumbs = append(context.Breadcrumbs, "Меню")
	}

	switch context.Breadcrumbs[len(context.Breadcrumbs)-1] {
	case "ЛК Отходообразователя":
		context.MenuItems = []MenuItem{
			{
				ID:    1,
				Label: "Вывести места образования отходов",
				Action: func(context *MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "ЛК Отходообразователя")
				},
			},
			{
				ID:    2,
				Label: "Создать место образования отходов",
				Action: func(context *MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "ЛК Отходообразователя")
				},
			},
			{
				ID:    3,
				Label: "Удалить место образования отходов",
				Action: func(context *MenuContext) {
					context.Stop()
				},
			},
			{
				ID:    4,
				Label: "Выход",
				Action: func(context *MenuContext) {
					context.Stop()
				},
			},
		}
	case "ЛК Оператора Хранения":
		context.MenuItems = []MenuItem{
			{
				ID:    1,
				Label: "Вывести место хранения отходов",
				Action: func(context *MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "ЛК Отходообразователя")
				},
			},
			{
				ID:    2,
				Label: "Создать место хранения отходов",
				Action: func(context *MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "ЛК Отходообразователя")
				},
			},
			{
				ID:    3,
				Label: "Удалить место хранения отходов",
				Action: func(context *MenuContext) {
					context.Stop()
				},
			},
			{
				ID:    4,
				Label: "Выход",
				Action: func(context *MenuContext) {
					context.Stop()
				},
			},
		}
	case "Симуляция":

	case "Меню":
		context.MenuItems = []MenuItem{
			{
				ID:    1,
				Label: "ЛК Отходообразователя",
				Action: func(context *MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "ЛК Отходообразователя")
				},
			},
			{
				ID:    2,
				Label: "ЛК Отходообразователя",
				Action: func(context *MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "ЛК Отходообразователя")
				},
			},
			{
				ID:    3,
				Label: "Симуляция",
				Action: func(context *MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "Симуляция")
				},
			},
			{
				ID:    4,
				Label: "Выход",
				Action: func(context *MenuContext) {
					context.Stop()
				},
			},
		}
	default:

	}
}

// Run запускает консольный интерфейс
func Run() {
	context := &MenuContext{
		isRunning: true,
	}
	getMenuContext(context)
	for context.isRunning {
		getMenuContext(context)
		displayMenu(context)

		choice := getUserChoice()

		if choice >= 0 && choice < len(context.MenuItems) {
			context.MenuItems[choice+1].Action(context)
		} else {
			fmt.Println("Неверный выбор. Пожалуйста, выберите опцию из списка.")
		}
	}
}
