package main

import "fmt"

func main() {
	var i, j, t int

	var m1 = make([][]int, 3, 3)
	var m2 = make([][]int, 3, 3)
	var ans = make([][]int, 3, 3)

	fmt.Println("Enter the First Matrix ")
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Scan(&t)
			m1[i] = append(m1[i], t)
		}
	}

	fmt.Println("Enter the Second Matrix ")
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Scan(&t)
			m1[i] = append(m1[i], t)
		}
	}
	add(m1, m2)
	ans = add(m1, m2)
	fmt.Println("Add Matrix ")
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Printf("%d  ", ans[i][j])
		}
		println()
	}
	ans = sub(m1, m2)
	fmt.Println("Sub Matrix ")
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Printf("%d  ", ans[i][j])
		}
		println()
	}
	ans = mul(m1, m2)
	fmt.Println("Mul Matrix ")
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Printf("%d  ", ans[i][j])
		}
		println()
	}

}
func add(mat1, mat2 [][]int) [][]int {
	mat1 = make([][]int, 3, 3)
	mat2 = make([][]int, 3, 3)
	var mat3 = make([][]int, 3, 3)
	var i, j, t int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			t = mat1[i][j] + mat2[i][j]
			mat3[i] = append(mat3[i], t)
		}
	}
	return mat3
}
func sub(mat1, mat2 [][]int) [][]int {
	mat1 = make([][]int, 3, 3)
	mat2 = make([][]int, 3, 3)
	var mat3 = make([][]int, 3, 3)
	var i, j, t int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			t = mat1[i][j] - mat2[i][j]
			mat3[i] = append(mat3[i], t)
		}
	}
	return mat3
}
func mul(mat1, mat2 [][]int) [][]int {
	mat1 = make([][]int, 3, 3)
	mat2 = make([][]int, 3, 3)
	var mat3 = make([][]int, 3, 3)
	var i, j, k int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			var c = 0
			for k = 0; k < 3; k++ {
				c = c + mat1[i][k]*mat2[k][j]
			}
			mat3[i] = append(mat3[i], c)

		}
	}
	return mat3
}
