package math

import (
	"fmt"
	"testing"
)

var testdata = []struct {
	FirstNumber  int
	SecondNumber int
	Expected     int
}{
	{2, 2, 4}, {3, 3, 9}, {10, 10, 100}, {3, 10, 30},
}

func TestMultiplyFunction(t *testing.T) {
	output := MultiplyTwoNumbers(20, 20)
	expected := 400
	if expected != int(output) {
		fmt.Println("Expected outut is: ", expected, "Output Recieved is: ", output)
	}
}

func TestBulkCheckOFMultiplyFunction(t *testing.T) {
	for i := 0; i < len(testdata); i++ {
		actual := MultiplyTwoNumbers(testdata[i].FirstNumber, testdata[i].SecondNumber)
		if actual != testdata[i].Expected {
			fmt.Println("Actual Output is :", actual, "  Expected Output is: ", testdata[i].Expected)
		}
	}
}
