package main

import "fmt"

type header struct {
	category string
	quantity int
	unitCost float64
}

var categories = []string{"household", "food", "drinks"}

var m = make(map[string]header)

func entireShoppingList() {

	m["Cups"] = header{category: "household", quantity: 5, unitCost: 3.0}
	m["Cake"] = header{category: "food", quantity: 3, unitCost: 1.0}
	m["Sprite"] = header{category: "drinks", quantity: 5, unitCost: 2.0}
	m["Fork"] = header{category: "household", quantity: 4, unitCost: 3.0}
	m["Bread"] = header{category: "food", quantity: 2, unitCost: 2.0}
	m["Plates"] = header{category: "household", quantity: 4, unitCost: 3.0}
	m["Coke"] = header{category: "drinks", quantity: 5, unitCost: 2.0}

}

func main() {
	entireShoppingList()
	for {
		var choiceMain int

		fmt.Println("")
		fmt.Println("Shopping list Application")
		fmt.Println("=========================")
		fmt.Println("1.View entire Shopping List report")
		fmt.Println("2.Generate Shopping List report")
		fmt.Println("3.Add items")
		fmt.Println("4.Modify items")
		fmt.Println("5.Delete items")
		fmt.Println("6.Print current data")
		fmt.Println("7.Add new category name")
		fmt.Println("8.End the program")
		fmt.Println("Select your Choice by typing 1 - 8:")
		fmt.Scanln(&choiceMain)

		if choiceMain == 1 {
			choiceOne()
		} else if choiceMain == 2 {
			choiceTwo()
		} else if choiceMain == 3 {
			choiceThree()

		} else if choiceMain == 4 {
			choiceFour()
		} else if choiceMain == 5 {
			choiceFive()
		} else if choiceMain == 6 {
			choiceSix()
		} else if choiceMain == 7 {
			choiceSeven()
		} else if choiceMain == 8 {
			break
		} else {
			fmt.Println("Please input number 1 - 8")
		}
	}
}
