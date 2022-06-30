package factory

import (
	"fmt"
	"module/celing"
	"module/exaust"
	"module/table"
)

// type fact interface {
// 	name(name string)
// 	getname() string
// }

type Factory struct {
	fan string
}

func (f *Factory) Categotise(fanType string) celing.Fact {

	switch fanType {
	case "celing":
		return celing.CreateCeling(fanType)
	case "exaust":
		return exaust.CreateExaust(fanType)
	case "table":
		return table.CreateTable(fanType)

	default:
		fmt.Println("No type Found")
	}
}
