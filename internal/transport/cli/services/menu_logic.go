package cli_services

import (
	models "GoProject/internal/transport/cli/models"
)

func GetMenuContext(context *models.MenuContext) {
	if len(context.Breadcrumbs) == 0 {
		context.Breadcrumbs = append(context.Breadcrumbs, "Меню")
	}

	switch context.Breadcrumbs[len(context.Breadcrumbs)-1] {
	case "ЛК Отходообразователя":
		context.MenuItems = []models.MenuItem{
			{
				ID:    1,
				Label: "Вывести места образования отходов",
				Action: func(context *models.MenuContext) {
					context.Stop()
				},
			},
			{
				ID:    2,
				Label: "Создать место образования отходов",
				Action: func(context *models.MenuContext) {
					context.Stop()
				},
			},
			{
				ID:    3,
				Label: "Удалить место образования отходов",
				Action: func(context *models.MenuContext) {
					context.Stop()
				},
			},
			{
				ID:    0,
				Label: "Выход",
				Action: func(context *models.MenuContext) {
					context.Stop()
				},
			},
		}
	case "ЛК Оператора Хранения":
		context.MenuItems = []models.MenuItem{
			{
				ID:    1,
				Label: "Вывести место хранения отходов",
				Action: func(context *models.MenuContext) {
					context.Stop()

				},
			},
			{
				ID:    2,
				Label: "Создать место хранения отходов",
				Action: func(context *models.MenuContext) {
					context.Stop()
				},
			},
			{
				ID:    3,
				Label: "Удалить место хранения отходов",
				Action: func(context *models.MenuContext) {
					context.Stop()
				},
			},
			{
				ID:    0,
				Label: "Выход",
				Action: func(context *models.MenuContext) {
					context.Stop()
				},
			},
		}
	case "Симуляция":

	case "Меню":
		context.MenuItems = []models.MenuItem{
			{
				ID:    1,
				Label: "ЛК Отходообразователя",
				Action: func(context *models.MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "ЛК Отходообразователя")
				},
			},
			{
				ID:    2,
				Label: "ЛК Отходообразователя",
				Action: func(context *models.MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "ЛК Отходообразователя")
				},
			},
			{
				ID:    3,
				Label: "Симуляция",
				Action: func(context *models.MenuContext) {
					context.Breadcrumbs = append(context.Breadcrumbs, "Симуляция")
					context.Stop()
				},
			},
			{
				ID:    0,
				Label: "Выход",
				Action: func(context *models.MenuContext) {
					context.Stop()
				},
			},
		}
	default:

	}
}
