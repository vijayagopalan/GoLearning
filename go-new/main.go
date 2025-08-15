package main

import (
	"fmt"
	menus "new-app/menu"
)

func main() {
	flag := false
	for flag {
		var new_id int
		var new_name string
		var new_qty int
		var choice string
		fmt.Println("Please selecty a option")
		fmt.Println("1. Add a Item")
		fmt.Println("2. Remove a Item")
		fmt.Println("3. Update a item")
		fmt.Println("4. get the item")
		fmt.Println("5. Exit")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			fmt.Println("enter id")
			fmt.Scan(&new_id)
			fmt.Println("enter name")
			fmt.Scan(&new_name)
			fmt.Println("enter qty")
			fmt.Scan(&new_qty)
			menus.Items.Add(new_name, new_id, new_qty)
		case "2":
			fmt.Println("enter id")
			fmt.Scan(&new_id)
			// menus.DeleteItem(&menus.Items, new_id)
		case "3":
			fmt.Println("enter id")
			fmt.Scan(&new_id)
			fmt.Println("enter name")
			fmt.Scan(&new_name)
			fmt.Println("enter qty")
			fmt.Scan(&new_qty)
			menus.Items.Update(new_name, new_id, new_qty)
		case "4":
			menus.Items.Print()
		case "5":
			flag = false
		default:
			fmt.Println("erro. please elect right option")
		}
	}

}
