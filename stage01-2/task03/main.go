package main

import "fmt"

/*
面向对象
1、题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
2、题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加
EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/
func main() {
	// 题目1
	// var rectangle Rectangle
	// rectangle.Area()
	// rectangle.Perimeter()
	// var circle Circle
	// circle.Area()
	// circle.Perimeter()

	// 题目2
	person := Person{Name: "abc", Age: 18}
	employee := Employee{Person: person, EmployeeId: 1}
	employee.PrintInfo()
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}
type Circle struct {
}

func (r Rectangle) Area() {
	fmt.Println("Rectangle ----doing AreaMethod")
}

func (r Rectangle) Perimeter() {
	fmt.Println("Rectangle ----doing Perimeter")
}

func (r Circle) Area() {
	fmt.Println("Circle ----doing AreaMethod")
}

func (r Circle) Perimeter() {
	fmt.Println("Circle ----doing Perimeter")
}

/*
题目2

*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId int
}

func (employee Employee) PrintInfo() {
	fmt.Printf("employee.Person.Name: %v\n", employee.Person.Name)
	fmt.Printf("employee.Person.Age: %v\n", employee.Person.Age)
	fmt.Printf("employee.EmployeeId: %v\n", employee.EmployeeId)
}
