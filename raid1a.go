package main

import "fmt"

func Raid1a(rows int, columns int) {
	for i := 1; i <= columns; i++ {
		for j := 1; j <= rows; j++ { //
			if (i == 1 && j == rows) || (i == 1 && j == 1) ||
				(i == columns && j == 1) || (i == columns && j == rows) {
				fmt.Print("o")
			} else if i != 1 && j == 1 || i != columns && j == rows {
				fmt.Print("|")
			} else if i == 1 || j == 1 || i == columns || j == rows {
				fmt.Print("-")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()

	}
}
func main() {
	Raid1a(5, 3)
}
