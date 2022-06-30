package model

import "github.com/google/uuid"

type Student struct {
	SID      string
	Fname    string
	Lname    string
	Division string
	Marks    []int
	Roll     int
	permanentaddress address
	tempaddress address
}
type address struct{
	typeofaddress bool
	pincode int
	city string
	state string

}

func CreateNewStudent(Fname, Lname, Division string, Roll int) *Student {

	tempStudent := &Student{
		Fname:    Fname,
		Lname:    Lname,
		Division: Division,
		Roll:     Roll,
		SID:      generateUUID(),
	}
	return tempStudent
}
func generateUUID() string {
	generateUUID := uuid.New()
	return generateUUID.String()
}
func (std *Student) Fullname() string {
	return std.Fname + " " + std.Lname
}
func (std *Student) InitalizeMarks(marks []int) {
	for i := 0; i < len(marks); i++ {
		std.Marks = append(std.Marks, marks[i])
	}

}
func address
