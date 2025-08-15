package menus

import "fmt"

type menu struct {
	id   int
	name string
	qty  int
}

type menuItem []menu

func (m menuItem) Print() {
	for _, value := range m {
		fmt.Println("id : ", value.id, "name : ", value.name, "qty : ", value.qty)
	}
}

func (m *menuItem) Add(name string, id, qty int) {
	new_record := menu{
		id:   id,
		name: name,
		qty:  qty,
	}
	*m = append(*m, new_record)
	fmt.Println(m)
}

func (m *menuItem) Update(name string, id, qty int) {
	new_inv := []menu{}
	for _, val := range *m {
		if val.id == id {
			new_record := menu{
				id:   id,
				name: name,
				qty:  qty,
			}
			new_inv = append(new_inv, new_record)
		} else {
			new_inv = append(new_inv, val)
		}
	}
	*m = new_inv
}

func (m *menuItem) Delete(id int) {
	new_inv := []menu{}
	for _, val := range *m {
		if val.id != id {
			new_inv = append(new_inv, val)
		}
	}
	*m = new_inv
}

func AddItem(invt *[]menu, name string, id, qty int) {
	new_record := menu{
		id:   id,
		name: name,
		qty:  qty,
	}
	*invt = append(*invt, new_record)
	fmt.Println(*invt)
}

func UpdateItem(invt *[]menu, name string, id, qty int) {
	new_inv := []menu{}
	for _, val := range *invt {
		if val.id == id {
			new_record := menu{
				id:   id,
				name: name,
				qty:  qty,
			}
			new_inv = append(new_inv, new_record)
		} else {
			new_inv = append(new_inv, val)
		}
	}
	*invt = new_inv
}

func DeleteItem(invt *[]menu, id int) {
	new_inv := []menu{}
	for _, val := range *invt {
		if val.id != id {
			new_inv = append(new_inv, val)
		}
	}
	*invt = new_inv
}

func PrintItems(invt *[]menu) {
	for _, value := range *invt {
		fmt.Println("id : ", value.id, "name : ", value.name, "qty : ", value.qty)
	}
}
