package main

type FanFactory interface {
	setName(name string)
	getName(name string)
}
type Fan struct {
	name string
}

func (f *Fan) setName(name string) {
	f.name = name

}

func (f *Fan) getName() string {
	return f.name
}

type celing struct {
	Fan
}

func (f *Fan) NewCeling(TypeOfFan string) FanFactory {
	return &celing{
		Fan: Fan{
			name: TypeOfFan,
		},
	}

}
func GetCelling() {

}
