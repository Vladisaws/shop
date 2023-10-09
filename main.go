package main

import "fmt"

func main() {
	var personBalans int
	var addProduct string
	var cart []string
	var pricePayable int
	shopp := map[string]int{
		"Лапша":    80,
		"Сосиски":  250,
		"Пельмени": 350,
	}

	fmt.Println("Пополните баланс")
	fmt.Scan(&personBalans)
	fmt.Println("Выберите продукты")
	for product, price := range shopp {
		fmt.Printf("%s: %d рублей\n", product, price)
	}
	fmt.Printf("Добавьте продукты в корзину.\n Введите имя продукта:")

	for fmt.Scan(&addProduct); addProduct != "все"; fmt.Scan(&addProduct) {
		for product, price := range shopp {
			if addProduct == product {
				cart = append(cart, product)
				pricePayable += price
			}
		}
		fmt.Printf("В корзине: %s\nНа сумму: %d\n", cart, pricePayable)
		fmt.Println("Введите имя продукта или напишите \"все\", если желаете закончить покупку")
	}
	remainingBalance := personBalans - pricePayable
	fmt.Printf("Вы купили: %s\n Остаток на счете: %d", cart, remainingBalance)

}
