package main

import "fmt"

func pascalTriangle(numRows int) [][]int {
	matrix := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		for j := 1; j <= i; j++ {
			if j == i {
				row[j] = 1
			}

			if j < i {
				row[j] = matrix[i-1][j-1] + matrix[i-1][j]
			}

		}
		matrix[i] = row
	}
	return matrix
}

func main() {
	fmt.Println(pascalTriangle(1))
	fmt.Println(pascalTriangle(2))
	fmt.Println(pascalTriangle(3))
	fmt.Println(pascalTriangle(4))
	fmt.Println(pascalTriangle(5))
	fmt.Println(pascalTriangle(6))
	fmt.Println(pascalTriangle(30))
}
