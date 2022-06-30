package exaust

import (
	"fmt"
	"module/celing"
)
type Exaust struct {
	fan string
}

func CreateExaust(fan string) celing.Fact {
	// var exaust *Exaust
	// exaust.fan = fan
	return exaust {
		fanName: fan,
	}
}
func (e *Exaust) print() {
	fmt.Println(e.fan)
}
