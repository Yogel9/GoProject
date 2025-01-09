package main

import (
	"GoProject/utils"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) add_year_1() int {
	utils.Add(p.Age, 1)
	fmt.Println("In resiver func, Age:", p.Age)
	return p.Age
}

func (p *Person) add_year_2() int {
	p.Age++
	return p.Age
}

func (p Person) struct_to_json() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	data, _ := json.Marshal(map[string]string{p.Name: string(p.Age)})
	fmt.Println(string(data))
}

func main() {
	p := Person{Name: "Anton", Age: 24}

	fmt.Println("Hello, World!")

	fmt.Println(p)
	fmt.Println(p.add_year_1())
	fmt.Println(p)
	fmt.Println(p.add_year_2())
	p.struct_to_json()
}

///9. Модули и зависимости
///10. Горутины и каналы
