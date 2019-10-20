package student

import "github.com/01-edu/z01"

func Raid1c(rows int, columns int) {
	a := 'A'
	b := 'A'
	c := 'C'
	d := 'C'
	wallg := 'B'
	wallv := 'B'
	if rows > 0 && columns > 0 {
		for i := 1; i <= columns; i++ {
			for j := 1; j <= rows; j++ {
				if i == 1 && j == 1 {
					z01.PrintRune(a)
				} else if i == 1 && j == rows {
					z01.PrintRune(b)
				} else if i == columns && j == 1 {
					z01.PrintRune(c)
				} else if i == columns && j == rows {
					z01.PrintRune(d) // o
				} else if i != 1 && j == 1 || i != columns && j == rows {
					z01.PrintRune(wallv) // |
				} else if i == 1 || j == 1 || i == columns || j == rows {
					z01.PrintRune(wallg) // -
				} else {
					z01.PrintRune(' ') // _
				}

			}
			z01.PrintRune(10)
		}

	}

}
