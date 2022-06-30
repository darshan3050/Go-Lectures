package view

type User struct {
	Unique   string
	Fname    string
	Lname    string
	Age      int
	Location string
	Number   int
	IsActive bool
}

var AllUser []User

func NewUser(fname, lname, location, uniqueName string, age, number int) *User {

	return &User{

		Unique:   uniqueName,
		Fname:    fname,
		Lname:    lname,
		Age:      age,
		Location: location,
		Number:   number,
		IsActive: true,
	}
}
