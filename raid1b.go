package student

import "github.com/01-edu/z01"

func Raid1b(rows int, columns int) {
	for i := 1; i <= columns; i++ {
		for j := 1; j <= rows; j++ {
			if (i == 1 && j == rows) || (i == columns && j == 1) {
				z01.PrintRune('\\') // a
			} else if (i == columns && j == rows) || (i == 1 && j == 1) {
				z01.PrintRune('/')
			} else if i != 1 && j == 1 || i != columns && j == rows {
				z01.PrintRune('*') // |
			} else if i == 1 || j == 1 || i == columns || j == rows {
				z01.PrintRune('*') // -
			} else {
				z01.PrintRune(' ') // _
			}

		}
		z01.PrintRune(10)
	}
}
