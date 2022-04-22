package main

import "fmt"

func choiceOne() {
	fmt.Println("1.View entire Shopping List report")

	for key, element := range m {
		fmt.Printf("Category: %v - Item: %v  Quantity: %v  Unit Cost: $%.2f\n", element.category, key, element.quantity, element.unitCost)
	}
}

func choiceTwo() {
	var choice int
	fmt.Println("2.Generate Shopping List report")
	fmt.Println("Generate report")
	fmt.Println("1. Total cost of each category")
	fmt.Println("2.List of item by category")
	fmt.Println("3.Main menu")
	fmt.Println("Choose your report")
	fmt.Scanln(&choice)

	switch choice {
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
}
func choiceThree() {
	fmt.Println("3.Add items")
	var newItem string
	var newCat string
	var newUnit int
	var newCost float64

	fmt.Println("What is the name of your new item? ")
	fmt.Scanln(&newItem)
	fmt.Println("What category does it belong to? ")
	fmt.Scanln(&newCat)
	fmt.Println("How many units? ")
	fmt.Scanln(&newUnit)
	fmt.Println("How much does it cost per unit? ")
	fmt.Scanln(&newCost)
	m[newItem] = header{newCat, newUnit, newCost}

	fmt.Printf("New updated items!!")
	for key, elements := range m {
		fmt.Printf("\nCategory: %v, Item: %v, Quantity:%v, Unit cost: $%.2f ", elements.category, key, elements.quantity, elements.unitCost)
	}
}
func choiceFour() {
	fmt.Println("4.Modify items (Only quantity change)")

	var newCategory, itemName string
	var newQuantity int
	var newCost float64

	fmt.Println("Which item would you like to modify?")
	fmt.Scanln(&itemName)
	// Coke

	fmt.Printf("Current item is Category: %v - Item: %v  Quantity: %v  Unit Cost: $%.2f\n", m[itemName].category, itemName, m[itemName].quantity, m[itemName].unitCost)

	// fmt.Println("Enter new name. Enter for no change.")
	// fmt.Scanln(&newName)
	fmt.Println("Enter new category. Enter for no change.")
	fmt.Scanln(&newCategory)

	if newCategory == "" {
		newCategory = m[itemName].category
	}

	fmt.Println("Enter new quantity. Enter for no change.")
	fmt.Scanln(&newQuantity)
	if newQuantity == 0 {
		newQuantity = m[itemName].quantity
	}
	// 7
	fmt.Println("Enter new unit cost. Enter for no change.")
	fmt.Scanln(&newCost)
	if newCost == 0.0 {
		newCost = m[itemName].unitCost
	}
	fmt.Println(&m)
	m[itemName] = header{newCategory, newQuantity, newCost}

	fmt.Println(m[itemName])

	// if newCategory == "" {
	// 	fmt.Println("No changes to category made.")
	// } else {
	// 	m[newCategory] = header{m[itemName].category, m[itemName].quantity, m[itemName].unitCost}
	// 	itemName = newCategory
	// }

	// if newQuantity == 0 {
	// 			fmt.Println("No changes to name made.")
	// 		} else {
	// 			m[newQuantity] = header{m[itemName].category, m[itemName].quantity, m[itemName].unitCost}
	// 			itemName = newName
	// 		}

	// 		if newCost == 0.0 {
	// 			fmt.Println("No changes to name made.")
	// 		} else {
	// 			m[newName] = header{m[itemName].category, m[itemName].quantity, m[itemName].unitCost}
	// 			itemName = newName
	// 		}
}
func choiceFive() {
	fmt.Println("5.Delete items")

	var itemDelete string
	fmt.Printf("Enter item name to delete: ")
	fmt.Scanln(&itemDelete)

	if _, ok := m[itemDelete]; ok {
		delete(m, itemDelete)
		fmt.Printf("Deleted %v\n", itemDelete)
	} else {
		fmt.Println("Item not found. Nothing to delete!")
	}
	fmt.Println(m)
	// if itemDelete := mw {
	//   delete(m, itemDelete)
	//   fmt.Printf("Deleted %v\n", itemDelete)
	// } else {
	//   fmt.Println("Item not found. Nothing to delete!")
	// }
}
func choiceSix() {
	fmt.Println("6.Print current data")

	fmt.Println(m)
}

func choiceSeven() {
	fmt.Println("7.Add new category name")
	var newCategory string
	duplicate := false
	fmt.Scanln(&newCategory)

	for _, category := range categories {
		if category == newCategory {
			duplicate = true
			continue
		}
	}

	if duplicate {
		fmt.Printf("already exists")
	} else {
		categories = append(categories, newCategory)
		fmt.Printf("%v added at index %v", newCategory, len(categories))
	}
}
