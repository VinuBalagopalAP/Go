package main

import (
	"fmt"
)

func main() {

	var row, col int

	var mat [10][10]string

	fmt.Print("Enter no of rows: ")

	fmt.Scanln(&row)

	fmt.Print("Enter no of column: ")

	fmt.Scanln(&col)

	fmt.Println("\nEnter matrix elements: ")

	for i := 0; i < row; i++ {

		for j := 0; j < col; j++ {

			fmt.Scanf("%s ", &mat[i][j])

		}

	}

	fmt.Println("\nMatrix is: \n")

	for i := 0; i < row; i++ {

		for j := 0; j < col; j++ {

			fmt.Printf("%s\t\t", mat[i][j])

		}
		fmt.Println("\n")

	}

}
