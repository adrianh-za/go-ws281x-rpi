package utils

import (
	"fmt"
	"os"
	"strings"
)

// Check to see if -v command flag passed in, if so then fmt.Println
// This was added as on RPi Zero, printing out to console causes flicker on the LED hat
func VerbosePrintln(a ...interface{}) (n int, err error) {
	for count := 0; count < len(os.Args); count++ {
		if strings.ToLower(os.Args[count]) == "-v" {
			return fmt.Println(a)
		}
	}
	return 0, nil
}

//Given the rows and columns of a LED hat, this will return sequential LED positions of the border of the hat
func GetLedOuterBoundry(rows int, columns int, reverse bool) []int {
	var totalLEDs = rows * columns
	var ledIndexes = make([]int, 0)

	//First row
	for count := 0; count < columns; count++ {
		ledIndexes = append(ledIndexes, count)
	}
	//First column
	for count := 2; count < rows; count++ {
		ledIndexes = append(ledIndexes, (count*columns)-1)
	}
	//Second row (opposite direction)
	for count := totalLEDs; count > (totalLEDs - columns); count-- {
		ledIndexes = append(ledIndexes, count-1)
	}
	//Second column (opposite direction)
	for count := totalLEDs - (columns * 2); count > 0; count = count - columns {
		ledIndexes = append(ledIndexes, count)
	}

	//Return the array of INT, reverse is specified
	if !reverse {
		return ledIndexes
	} else {
		for i, j := 0, len(ledIndexes)-1; i < j; i, j = i+1, j-1 {
			ledIndexes[i], ledIndexes[j] = ledIndexes[j], ledIndexes[i]
		}
		return ledIndexes
	}
}