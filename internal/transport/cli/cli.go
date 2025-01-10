package cli

import (
	models "GoProject/internal/transport/cli/models"
	services "GoProject/internal/transport/cli/services"
	"fmt"
)

func Run() {
	context := &models.MenuContext{}
	context.Start()
	services.GetMenuContext(context)
	for context.IsActive() {
		services.GetMenuContext(context)
		services.DisplayMenu(context)

		choice := services.GetUserChoice()

		if choice >= 0 && choice < len(context.MenuItems) {
			context.RunAction(choice)
		} else {
			fmt.Println("Неверный выбор. Пожалуйста, выберите опцию из списка.")
		}
	}
}
