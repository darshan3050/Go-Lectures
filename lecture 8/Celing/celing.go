package celing

import "fmt"

type Fact interface {
	// name(name string)
	// getname() string
	print()
}

type Celing struct {
	fanName string
}

func CreateCeling(fan string) Fact{
	return &Celing{
		fanName: fan,
	}
}
func (c *Celing) print()  {
	fmt.Println(c.fanName)
}
