package day5

import (
	"errors"
	"fmt"
	"log"
	"math"
)

//					   64    16    2     1
// row_upperbound: 127 -> 63 -> 47 -> 45 -> 44
//					   32 	 8     4
// row_lowerbound: 0   -> 32 -> 40 -> 44
func binaryBoard( directions string, planeRows, planeCol int)int {
	col_lowerBound := 0
	col_upperbound := planeCol - 1
	row_lowerbound := 0
	row_upperbound := planeRows - 1
	rowhalf := planeRows
	colhalf := planeCol
	for i := 0; i < len(directions); i++ {
		if (col_lowerBound > col_upperbound) || (row_lowerbound > row_upperbound) {
			fmt.Printf("col_lowerBound: %d, col_upperbound: %d\n",col_lowerBound, col_upperbound)
			fmt.Printf("row_lowerBound: %d, row_upperbound: %d\n",row_lowerbound, row_upperbound)
			log.Fatal(errors.New("invalid directions."))
		}
		if directions[i] == 'F' {
			rowhalf = rowhalf /2
			row_upperbound = row_upperbound - rowhalf
		} else if directions[i] == 'B' {
			rowhalf = rowhalf /2
			row_lowerbound = row_lowerbound + rowhalf
		} else if directions[i] == 'L' {
			colhalf = colhalf /2
			col_upperbound = col_upperbound - colhalf

		} else if directions[i] == 'R' {
			colhalf = colhalf /2
			col_lowerBound = col_lowerBound + colhalf
		}
	}
	row := int(math.Min(float64(row_lowerbound), float64(row_upperbound)))
	col := int(math.Max(float64(col_lowerBound), float64(col_upperbound)))
	return row * planeCol + col
}

// exactly one of the 128 rows on the plane (numbered 0 through 127)
// exactly one of the 8 columns of seats on the plane (numbered 0 through 7).
func binaryBoardFixedSize(directions string) int {
	return binaryBoard(directions, 128, 8)
}