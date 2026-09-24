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

// Возможные скидки
const (
	birtdayDiscount float64 = 16
	studentDiscount float64 = 12.8
	seniorDiscount float64 = 25.6
	simpleDiscount float64 = 51.2
)


// Глобальная переменная для подсчета товаров
var totalItems int

// Глобальный массив товаров
var arr [MaxItems]Item

// Структура предмета
type Item struct {
	itemName string
	itemPrice float64
	minQuantity, maxQuantity int
	isAvailable bool
	quantity int
	itemWeight float32
	percentInStock float64
	category string
	itemColor string
	itemID int64
	itemProducer string
	dateAdded time.Time
}

func main() {
	// Создание предмета из шаблона просто в мэйне
	var itemName string
	itemName = "Смартфон"
	var itemPrice float64 = 999.99
	var minQuantity, maxQuantity int = 5, 20
	var isAvailable = true
	quantity := 15
	var (
		itemID     int64  = 12345
		itemColor  string = "Черный"
		itemWeight float32
	)
	itemWeight = 0.3
	category := CategoryElectronics

	addNewItem(itemID, itemName, quantity, itemPrice, isAvailable, category, itemWeight, minQuantity, maxQuantity, itemColor)
	displayItemInfo(arr[0])

	fmt.Printf("\nОбщее количество товаров на складе: %d\n", totalItems)
	// addNewItem()

	displayItemExtraInfo(arr[0])

	//Для проверки наличия чего-то типо бдшки
	fmt.Println("\n", arr[0], arr[1])
	calculateDiscount(arr[0].itemPrice)
}

// Функция для отображения информации о товаре
func displayItemInfo(item Item) {
	fmt.Println("=== Информация о товаре ===")
	fmt.Printf("ID: %d\n", item.itemID)
	fmt.Printf("Название: %s\n", item.itemName)
	fmt.Printf("Категория: %s\n", item.category)
	fmt.Printf("Количество: %d\n", item.quantity)
	fmt.Printf("Цена: %.2f руб.\n", item.itemPrice)

	// Использование условного оператора с логическим типом
	if item.isAvailable {
		fmt.Println("Статус: В наличии")
	} else {
		fmt.Println("Статус: Нет в наличии")
	}
}

// Доп инфа (не очень понимаю для чего, но решил перенести сюда)
func extraInfo() {

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
func addNewItem(id int64, name string, qty int, price float64, isAvailable bool, category string, itemWeight float32, minQ int, maxQ int, color string) int64 {
	var newItem Item
	newItem.itemID = id
	newItem.itemName = name
	newItem.itemPrice = price
	newItem.minQuantity = minQ
	newItem.maxQuantity = maxQ
	newItem.isAvailable = isAvailable
	newItem.quantity = qty
	newItem.itemWeight = itemWeight
	newItem.percentInStock = float64(qty) / float64(MaxItems) * 100
	newItem.category = category
	newItem.itemColor = color
	newItem.dateAdded = time.Now()

	for i := 0; i<newItem.quantity; i++ {
		arr[totalItems + i] = newItem
	}
	updateTotalItems(qty)
	return arr[totalItems].itemID
}

// Функция отображения доп информации
func displayItemExtraInfo(item Item){
	switch item.category {
	case CategoryElectronics:
		fmt.Println("\nДополнительная информация:")
		fmt.Printf("Цвет товара: %s\n", item.itemColor)
		fmt.Printf("Вес товара: %.2f кг\n", item.itemWeight)
		fmt.Printf("Минимальное количество: %d, Максимальное количество: %d\n", item.minQuantity, item.maxQuantity)
		fmt.Printf("Процент от максимального количества на складе: %.1f%%\n", item.percentInStock)
		fmt.Printf("Дата добавления: %s\n", item.dateAdded.Format("02-01-2006"))
	case CategoryClothes:
		fmt.Println("\nДополнительная информация:")
		fmt.Printf("Цвет товара: %s\n", item.itemColor)
		fmt.Printf("Вес товара: %.2f кг\n", item.itemWeight)
		fmt.Printf("Минимальное количество: %d, Максимальное количество: %d\n", item.minQuantity, item.maxQuantity)
		fmt.Printf("Процент от максимального количества на складе: %.1f%%\n", item.percentInStock)
		fmt.Printf("Дата добавления: %s\n", item.dateAdded.Format("02-01-2006"))
	case CategoryFood:
		fmt.Println("\nДополнительная информация:")
		fmt.Printf("Вес товара: %.2f кг\n", item.itemWeight)
		fmt.Printf("Минимальное количество: %d, Максимальное количество: %d\n", item.minQuantity, item.maxQuantity)
		fmt.Printf("Процент от максимального количества на складе: %.1f%%\n", item.percentInStock)
		fmt.Printf("Дата добавления: %s\n", item.dateAdded.Format("02-01-2006"))
	default:
		println("Error: у этого предмета нет категории")
	}
}

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