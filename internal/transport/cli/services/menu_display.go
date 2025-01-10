package cli_services

import (
	models "GoProject/internal/transport/cli/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для отображения меню
func DisplayMenu(context *models.MenuContext) {
	fmt.Println("=== Контекстное меню ===")
	fmt.Printf("Path: %v\n", strings.Join(context.Breadcrumbs, " >> "))
	for _, item := range context.MenuItems {
		fmt.Printf("[%d]. %s\n", item.ID, item.Label)
	}
	fmt.Print("Выберите опцию: ")
}

// Функция для получения выбора пользователя
func GetUserChoice() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		return -1 // Возвращаем -1, если ввод некорректен
	}
	return choice
}
