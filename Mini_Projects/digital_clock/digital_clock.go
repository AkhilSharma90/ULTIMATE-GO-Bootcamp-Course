package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// Grid
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	// store 10 digital array in `digit`
	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	// Let terminal screen clear. By the package `github.com/inancgumus/screen`  (Cannot run in go playground)

	screen.Clear()

	// infinite loop to show in terminal forever.
	// Don't run this in Go playground. Run this instead => for i :=0 i<1000; i++
	for {

		// Move the cursol in the top-left. To put the clock always in the same position.
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		// fmt.Printf("hour: %d min: %d sec:%d\n", hour, min, sec)

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		// loop five (grid) times
		for line := 0; line < 5; line++ {
			for index, digit := range clock {
				//like:	array	 zero	   0
				next := clock[index][line]

				// if the clock array's digit is colon array, then use blank to skip to showing.
				// hide colon in each 2 sec
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				// 2 space between number
				fmt.Print(next, "  ")
			}
			// to make new line after make 1 row(grid)
			fmt.Println()

		}
		fmt.Println()
		// update by each sec.
		time.Sleep(time.Second)
	}
}

// If shows `1 2 3 4 5 6 7 8 9` in terminal, use below instead of `clock`
// // digits[0] is the first digit array (zero) but could used second digit array as well
// // same as: for line := range digits[0] {
// // loop five (grid) times
// for line := 0; line < 5; line++ {
// 	for digit := range digits {
// 		//like		array		zero	 0		2 space between number
// 		fmt.Print(digits[digit][line], "  ")
// 	}
// 	// to make new line after make 1 row(grid)
// 	fmt.Println()
// }
