package piscine

import "github.com/01-edu/z01"

func Raid1a(rows int, columns int) {
	for i := 1; i <= columns; i++ {
		for j := 1; j <= rows; j++ { //
			if (i == 1 && j == rows) || (i == 1 && j == 1) ||
				(i == columns && j == 1) || (i == columns && j == rows) {
				z01.PrintRune(111)
			} else if i != 1 && j == 1 || i != columns && j == rows {
				z01.PrintRune(124)
			} else if i == 1 || j == 1 || i == columns || j == rows {
				z01.PrintRune(45)
			} else {
				z01.PrintRune(32)
			}

		}
		z01.PrintRune(10)

	}
}
