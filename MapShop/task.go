package main

import "fmt"

func main() {
	shopList := make(map[string]int)

	shopList["хлеб"] = 50
	shopList["молоко"] = 100
	shopList["масло"] = 200
	shopList["колбаса"] = 500
	shopList["соль"] = 20
	shopList["огурцы"] = 200
	shopList["сыр"] = 600
	shopList["ветчина"] = 700
	shopList["буженина"] = 900
	shopList["помидоры"] = 250
	shopList["рыба"] = 300
	shopList["хамон"] = 1500

	price := 500
	fmt.Println("Order price above ", price, ": ", getGoodsAbovePrice(shopList, price))
	fmt.Println()

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	fmt.Println("Order: ", order)
	fmt.Println("Order price: ", calcOrderPrice(shopList, order))
}

func getGoodsAbovePrice(shopList map[string]int, priceAbove int) []string {
	var result []string
	result = make([]string, 0, 12)

	for goods, price := range shopList {
		if price > priceAbove {
			result = append(result, goods)
		}
	}

	return result
}

func calcOrderPrice(shopList map[string]int, orderList []string) int {
	var result int

	for _, order := range orderList {
		result += shopList[order]
	}

	return result
}
