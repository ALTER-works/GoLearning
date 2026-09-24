// Программа для управления небольшим складом товаров
package main

import (
	"fmt"
	"time"
)

// Глобальные константы для категорий товаров
const (
	CategoryElectronics = "Электроника"
	CategoryFood        = "Продукты"
	CategoryClothes     = "Одежда"
	MaxItems            = 100 // Максимальное количество товаров на складе
)

// Глобальная переменная для подсчета товаров
var totalItems int

func main() {
	// Объявление и инициализация переменных разными способами

	// Способ 1: Объявление переменной с явным указанием типа и последующим присваиванием
	var itemName string
	itemName = "Смартфон"

	// Способ 2: Объявление переменной с явным указанием типа и инициализацией
	var itemPrice float64 = 999.99

	// Способ 3: Объявление нескольких переменных одного типа
	var minQuantity, maxQuantity int = 5, 20

	// Способ 4: Объявление переменной без явного указания типа (тип определяется компилятором)
	var isAvailable = true

	// Способ 5: Короткое объявление переменной (только внутри функций)
	quantity := 15

	// Способ 6: Объявление нескольких переменных разных типов одновременно
	var (
		itemID     int64  = 12345
		itemColor  string = "Черный"
		itemWeight float32
		dateAdded  time.Time = time.Now()
	)

	// Присваивание значения ранее объявленной переменной
	itemWeight = 0.3

	// Приведение типов
	percentInStock := float64(quantity) / float64(MaxItems) * 100

	// Использование констант
	category := CategoryElectronics

	// Вызов функции для вывода информации о товаре
	displayItemInfo(itemID, itemName, quantity, itemPrice, isAvailable, category)

	// Демонстрация работы с базовыми типами
	fmt.Println("\nДополнительная информация:")
	fmt.Printf("Цвет товара: %s\n", itemColor)
	fmt.Printf("Вес товара: %.2f кг\n", itemWeight)
	fmt.Printf("Минимальное количество: %d, Максимальное количество: %d\n", minQuantity, maxQuantity)
	fmt.Printf("Процент от максимального количества на складе: %.1f%%\n", percentInStock)
	fmt.Printf("Дата добавления: %s\n", dateAdded.Format("02-01-2006"))

	// Работа с целочисленными константами и их типами
	const (
		Small  = 1
		Medium = 5
		Large  = 10
	)

	// Демонстрация работы с различными числовыми типами
	var (
		shortValue   int8      = 127 // -128 до 127
		intValue     int       = 1000000
		uintValue    uint      = 10000 // только положительные
		floatValue   float32   = 123.456
		complexValue complex64 = 1 + 2i
	)

	fmt.Println("\nРазные числовые типы:")
	fmt.Printf("int8: %d\n", shortValue)
	fmt.Printf("int: %d\n", intValue)
	fmt.Printf("uint: %d\n", uintValue)
	fmt.Printf("float32: %f\n", floatValue)
	fmt.Printf("complex64: %v\n", complexValue)

	// Работа со строками и байтами
	productCode := "ЯЩИК-12345"
	fmt.Printf("\nКод товара: %s, длина: %d символов, длина: %d байт\n", productCode, len([]rune(productCode)), len(productCode))

	// Обновление общего количества товаров
	updateTotalItems(quantity)
	fmt.Printf("\nОбщее количество товаров на складе: %d\n", totalItems)
	
	addNewItem()
	fmt.Printf("\nОбщее количество товаров на складе: %d\n", totalItems)
}

// Функция для отображения информации о товаре
func displayItemInfo(id int64, name string, qty int, price float64, isAvailable bool, category string) {
	fmt.Println("=== Информация о товаре ===")
	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Название: %s\n", name)
	fmt.Printf("Категория: %s\n", category)
	fmt.Printf("Количество: %d\n", qty)
	fmt.Printf("Цена: %.2f руб.\n", price)

	// Использование условного оператора с логическим типом
	if isAvailable {
		fmt.Println("Статус: В наличии")
	} else {
		fmt.Println("Статус: Нет в наличии")
	}
}

// Функция для обновления общего количества товаров
func updateTotalItems(qty int) {
	totalItems += qty

	// Проверка на превышение максимального количества
	if totalItems > MaxItems {
		fmt.Println("Предупреждение: Превышено максимальное количество товаров на складе!")
		totalItems = MaxItems
	}
}

// Функция добавления нового предмета
func addNewItem(){
	quantity := 15
	var itemID int64 = 123456
	var itemName = "Шляпа с перьями"
	var itemPrice float64 = 1449.99
	// var minQuantity = 10
	// var maxQuantity = 30 т.к. не использую, коменчу, ибо компилятор жалуется
	var itemIsAvailable = true
	itemCategory := CategoryClothes
	updateTotalItems(quantity)

	fmt.Printf("\nНовый предмет добавлен в хранилище. ID:%d \n\n", itemID)
	displayItemInfo(itemID, itemName, quantity, itemPrice, itemIsAvailable, itemCategory)

	fmt.Print("\nВозможные скидки: ")
	calculateDiscount(itemPrice)
}

// Возможные скидки
const (
	birtdayDiscount float64 = 16
	studentDiscount float64 = 12.8
	seniorDiscount float64 = 25.6
	simpleDiscount float64 = 51.2
)

// Функция рассчёта скидки
func calculateDiscount(price float64){
	studentD := price/100 * (100 - studentDiscount)
	birtdayD := price/100 * (100 - birtdayDiscount)
	seniorD := price/100 * (100 - seniorDiscount)
	simpleD := price/100 * (100 - simpleDiscount)

	fmt.Printf(`
* Цена со скидкой студенту: %3f
* Цена со скидкой в честь дня рождения: %3f
* Цена со скидкой пенсионеру: %3f
* Цена со скидкой знающему: %3f
	`, studentD, birtdayD, seniorD, simpleD)
}