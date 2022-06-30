package table

import "fmt"

type Table struct {
	fan string
}

func CreateTable(fan string) {
	var table *Table
	table.fan = fan
}
func (t *Table) printtable() {
	fmt.Println(t.fan)
}
