package main

import (
	"fmt";
)

const ProgName string = "GoFit"

// Объект юзера (класс)
type User struct {
	name string
	age int
	height float32
	have_subscr bool
}
// Метод объекта юзера, выводит всё инфо
func (someUser User) show_user() {
	fmt.Printf(`
Добро пожаловать в %s!
Профиль пользователя:
Имя: %s
Возраст: %d лет
Рост: %.3f м
Подписан на рассылку: %t
`, ProgName, someUser.name, someUser.age, 
someUser.height, someUser.have_subscr)
}

func get_user_info() User {
	var info User
	fmt.Scan(&info.name,&info.age,&info.height,&info.have_subscr)
	println(info.age, info.name)
	return info
}

func main(){
	var firstUser = User{"Arnold", 30, 1.76, false} 
	firstUser.show_user()

	var secondUser = get_user_info()
	secondUser.show_user()
}