package main

import "fmt"

// Структура для хранения данных о сотруднике
type Employee struct {
	name string
	sirName string
	age int
	position string
	salary int
}

// Интерфейс для вывода информации
type Displayable interface {
	Display()
}

// Реализация метода Display для Employee
func (e Employee) Display() {
	fmt.Printf(`
Информация о сотруднике:
Имя: %s
Фамилия: %s
Возраст: %d
Должность: %s
ЗП: %d
`, e.name, e.sirName,
e.age, e.position, e.salary,
)
}

// Функция фильтрации сотрудников по возрасту и зарплате
func FilterEmployees(employees []Employee, minAge int, minSalary int) []Employee {
	filtered := []Employee{}
	for _, employee := range employees{
		if (employee.age >= minAge) && (employee.salary >= minSalary){
			filtered = append(filtered, employee)
		} 
	}
	return filtered
}

func AllEmployeesInfo(employees []Employee){
	fmt.Printf(`
=================================
|	ИНФОРМАЦИЯ О СОТРУДНИКАХ	|
=================================`)

	for _, employee := range employees{
		employee.Display()
	}
}

func main() {
	// Инициализация списка сотрудников
	employees := []Employee{
		{
			name: "Анатолий",
			sirName: "Попов",
			age: 75,
			position: "Chief Technology Officer",
			salary: 500000,
		},
		{
			name: "Василий",
			sirName: "Пупкин",
			age: 20,
			position: "Internal Technical Support",
			salary: 30000,
		},
		{
			name: "Георгий",
			sirName: "Васильев",
			age: 31,
			position: "IT Support Engineer",
			salary: 200000,
		},
		{
			name: "Дмитрий",
			sirName: "Иванов",
			age: 50,
			position: "IT Support Engineer",
			salary: 200000,
		},
		{
			name: "Анна",
			sirName: "Павловна",
			age: 47,
			position: "Systems Administrator",
			salary: 180000,
		},
		{
			name: "Валентин",
			sirName: "Петрович",
			age: 82,
			position: "Cleaning Supervisor",
			salary: 80000,
		},
	}

	// Параметры фильтрации
	minAge :=  50
	minSalary := 100000

	// Фильтрация и вывод
	filteredEmployees := FilterEmployees(employees, minAge, minSalary)
	AllEmployeesInfo(filteredEmployees)
}