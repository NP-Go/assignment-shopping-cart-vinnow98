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
	m["Sprite"] = header{category: "Drinks", quantity: 5, unitCost: 2.0}
	m["Fork"] = header{category: "household", quantity: 4, unitCost: 3.0}
	m["Bread"] = header{category: "food", quantity: 2, unitCost: 2.0}
	m["Plates"] = header{category: "household", quantity: 4, unitCost: 3.0}
	m["Coke"] = header{category: "drinks", quantity: 5, unitCost: 2.0}

}

func main() {

	var choiceMain int
	var choice2 int
	fmt.Println("Shopping list Application")
	fmt.Println("=========================")
	fmt.Println("1.View entire Shopping List report")
	fmt.Println("2.Generate Shopping List report")
	fmt.Println("3.Add items")
	fmt.Println("4.Modify items")
	fmt.Println("5.Delete items")
	fmt.Println("6.Print current data")
	fmt.Println("7.Add new category name")
	fmt.Println("Select your Choice by typing 1 - 7:")
	fmt.Scanln(&choiceMain)

	entireShoppingList()
	fmt.Print(m["Cups"].quantity)
	if choiceMain == 1 {
		fmt.Print("1.View entire Shopping List report")

		for key, element := range m {
			fmt.Printf("Category: %v - Item: %v  Quantity: %v  Unit Cost: $%.2f\n", element.category, key, element.quantity, element.unitCost)
		}
	} else if choiceMain == 2 {
		fmt.Print("2.Generate Shopping List report")
		fmt.Println("Generate report")
		fmt.Println("1. Total cost of each category")
		fmt.Println("2.List of item by category")
		fmt.Println("3.Main menu")
		fmt.Println("Choose your report")
		fmt.Scanln(&choice2)

		switch choice2 {
		case 1:
			for _, elementOutside := range categories {
				sum := 0.00
				for _, elementInside := range m {
					if elementOutside == elementInside.category {
						sum = sum + (float64(elementInside.quantity) * elementInside.unitCost)
					}
				}
				fmt.Printf("%v cost: $%.2f\n", elementOutside, sum)
			}

		case 2:
			fmt.Println("")
			fmt.Println("List of items by category")
			fmt.Println("=========================")
			for _, elementOutside := range categories {
				for key, elementInside := range m {
					if elementOutside == elementInside.category {
						fmt.Printf("Category: %v - Item: %v  Quantity: %v  Unit Cost: $%.2f\n", elementOutside, key, elementInside.quantity, elementInside.unitCost)
					}
				}
			}
		case 3:
			break
		}
	} else if choiceMain == 3 {
		fmt.Print("3.Add items")
		var newItem string
		var newCat string
		var newUnit int
		var newCost float64

		fmt.Print("What is the name of your new item? ")
		fmt.Scanln(&newItem)
		fmt.Print("What category does it belong to? ")
		fmt.Scanln(&newCat)
		fmt.Print("How many units? ")
		fmt.Scanln(&newUnit)
		fmt.Print("How much does it cost per unit? ")
		fmt.Scanln(&newCost)
		m[newItem] = header{newCat, newUnit, newCost}

		fmt.Printf("New updated items!!")
		for key, elements := range m {
			fmt.Printf("\nCategory: %v, Item: %v, Quantity:%v, Unit cost: $%.2f ", elements.category, key, elements.quantity, elements.unitCost)
		}

	} else if choiceMain == 4 {
		fmt.Println("4.Modify items (Only quantity change)")

		var newName, newCategory, itemName string
		var newQuantity int
		var newCost float64

		fmt.Println("Which item would you like to modify?")
		fmt.Scanln(&itemName)
		fmt.Printf("Current item is Category: %v - Item: %v  Quantity: %v  Unit Cost: $%.2f\n", m[itemName].category, itemName, m[itemName].quantity, m[itemName].unitCost)

		fmt.Println("Enter new name. Enter for no change.")
		fmt.Scanln(&newName)
		fmt.Println("Enter new category. Enter for no change.")
		fmt.Scanln(&newCategory)
		fmt.Println("Enter new quantity. Enter for no change.")
		fmt.Scanln(&newQuantity)
		fmt.Println("Enter new unit cost. Enter for no change.")
		fmt.Scanln(&newCost)

		if newName == "" {
			fmt.Println("No changes to name made.")
		} else {
			m[newName] = header{m[itemName].category, m[itemName].quantity, m[itemName].unitCost}
			itemName = newName
		}

		if newCategory == "" {
			fmt.Println("No changes to category made.")
		} else {
			m[newCategory] = header{m[itemName].category, m[itemName].quantity, m[itemName].unitCost}
			itemName = newCategory
		}

		/*if newQuantity == 0 {
				fmt.Println("No changes to name made.")
			} else {
				m[newQuantity] = header{m[itemName].category, m[itemName].quantity, m[itemName].unitCost}
				itemName = newName
			}

			if newCost == 0.0 {
				fmt.Println("No changes to name made.")
			} else {
				m[newName] = header{m[itemName].category, m[itemName].quantity, m[itemName].unitCost}
				itemName = newName
			}


		} else if choiceMain == 5{
						fmt.Println("5.Delete items")

						var itemDelete string
						fmt.Printf("Enter item name to delete: ")
				        fmt.Scanln(&itemDelete)

				        if itemDelete := m[item] {
				          delete(m, item)
				          fmt.Printf("Deleted %v\n", itemDelete)
				        } else {
				          fmt.Println("Item not found. Nothing to delete!")
				        }
		}else if choiceMain == 6{

		}
		}else if choiceMain == 7{

		}*/
	}

}
