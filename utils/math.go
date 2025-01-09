package utils

import "fmt"

///выполнить какую-то инициализацию при загрузке пакета
func init() {
	fmt.Println("Math package initialized")
}

func Add(a, b int) int {
	return a + b
}
