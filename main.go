package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) add_year_1() int {
	p.Age++
	fmt.Println("In resiver func, Age:", p.Age)
	return p.Age
}

func (p *Person) add_year_2() int {
	p.Age++
	return p.Age
}

func main() {
	p := Person{Name: "Anton", Age: 24}

	fmt.Println("Hello, World!")

	fmt.Println(p)
	fmt.Println(p.add_year_1())
	fmt.Println(p)
	fmt.Println(p.add_year_2())
	fmt.Println(p)
}

///9. Модули и зависимости
///10. Горутины и каналы
