package main

import "fmt"

// Структура для хранения данных о сотруднике
type Employee struct {
	name     string
	surName  string
	age      int
	position string
	salary   int
}

// Костыль для задания. Плюс тест возможностей языка.
type Stuff struct {
	emploees []Employee
}

func (s Stuff) Display(){
	fmt.Printf(`
=================================
|	ИНФОРМАЦИЯ О СОТРУДНИКАХ	|
=================================`)

	for _, emploee := range s.emploees {
		emploee.Display()
	}
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
`, e.name, e.surName,
		e.age, e.position, e.salary,
	)
}

// Функция фильтрации сотрудников по возрасту и зарплате
func FilterEmployees(employees []Employee, minAge int, minSalary int) []Employee {
	filtered := []Employee{}
	for _, employee := range employees {
		if (employee.age >= minAge) && (employee.salary >= minSalary) {
			filtered = append(filtered, employee)
		}
	}
	return filtered
}

func AllObjectInfo(objects []Displayable) {
	fmt.Printf(`
=================================
|	ИНФОРМАЦИЯ О СОТРУДНИКАХ	|
=================================`)

	for _, obj := range objects {
		obj.Display()
	}
}

func IsDisplayable(d any) (string, error){
	if _, ok := d.(Displayable); ok {
		return fmt.Sprintln("Да, это Displayable."), nil
	} else {
		return fmt.Sprintln("Нет, это не Displayable."), fmt.Errorf("Объект не является Displayable")
	}
}

func main() {
	// Инициализация списка сотрудников
	employees := []Employee{
		{
			name:     "Анатолий",
			surName:  "Попов",
			age:      75,
			position: "Chief Technology Officer",
			salary:   500000,
		},
		{
			name:     "Василий",
			surName:  "Пупкин",
			age:      20,
			position: "Internal Technical Support",
			salary:   30000,
		},
		{
			name:     "Георгий",
			surName:  "Васильев",
			age:      31,
			position: "IT Support Engineer",
			salary:   200000,
		},
		{
			name:     "Дмитрий",
			surName:  "Иванов",
			age:      50,
			position: "IT Support Engineer",
			salary:   200000,
		},
		{
			name:     "Анна",
			surName:  "Павловна",
			age:      47,
			position: "Systems Administrator",
			salary:   180000,
		},
		{
			name:     "Валентин",
			surName:  "Петрович",
			age:      82,
			position: "Cleaning Supervisor",
			salary:   80000,
		},
	}

	// Параметры фильтрации
	minAge := 50
	minSalary := 100000

	// Фильтрация и вывод
	var filteredEmployees Displayable = Stuff{FilterEmployees(employees, minAge, minSalary)}
	filteredEmployees.Display()

	// Доп вывод типа данных. Интересно, что без пакетов узнать, является ли filteredEmployees по типу Displayable через fmt type нельзя. 
	// Только своими функциями через ошибку можно это всё красиво расписать. Ну, спасибо что у интерфейса есть метод встроенный самовызова.
	// Хотя чисто технически, это не метод, а условная запись для компилятора, чтобы он вызвал assertE2I для проверки хэшей (маппингов) между двумя типами.
	fmt.Printf("\nТип данных объекта, к которому применилась функция Display: %T", filteredEmployees)
	text, err := IsDisplayable(filteredEmployees)
	if err != nil {
		fmt.Printf("\nЯвляется ли объект Dicplayable: %s", text)
	} else {
		fmt.Printf("\nЯвляется ли объект Dicplayable: %s", text)
	}
}
