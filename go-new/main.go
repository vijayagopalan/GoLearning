package main

import (
	"bytes"
	"fmt"
	menus "new-app/menu"
	"strings"
	"sync"
)

type printer interface {
	Print() string
}

type user struct {
	id   int
	name string
}

func (u user) Print() string {
	return ("user id is " + fmt.Sprint(u.id) + "and  name is " + u.name)
}

type record struct {
	name   string
	prices map[string]float32
}

func (mi record) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for i, v := range mi.prices {
		fmt.Println("index", i, "value", v)
	}
	return b.String()
}

func clone[V any](inp []V) []V {
	new_value := make([]V, len(inp))
	for ind, val := range inp {
		new_value[ind] = val
	}
	return new_value
}

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
	var p printer
	var result string
	p = user{id: 1, name: "abc"}
	result = p.Print()
	fmt.Print(result)
	p = record{name: "coffer", prices: map[string]float32{
		"itm1": 10,
		"itm2": 10,
		"itm3": 10,
		"itm4": 10,
	}}
	result = p.Print()
	fmt.Print(result)

	scores := []float32{
		10.0,
		20.0,
		30.0,
		40.0,
	}

	c := clone(scores)
	fmt.Print(c)
	fmt.Print("\n" + strings.Repeat("-", 10) + "\n")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Print("async call" + "\n")
		wg.Done()
	}()
	fmt.Print("sync call" + "\n")
	wg.Wait()
	wg.Add(1)

	ch := make(chan string)
	go func() {
		ch <- "hi"
	}()

	go func() {
		fmt.Print(<-ch)
		wg.Done()
	}()
	wg.Wait()
}
