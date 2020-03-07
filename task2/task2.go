package task2

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(N int, S string, T string) string {
	// write your code in Go 1.4
	m := make(map[string]int)

	sunk := 0
	survived := 0
	for i := 0; i < N; i++ {
		m[strconv.Itoa(i+1)] = i
		m[string('A'+i)] = i
	}

	hitMap := getHitMap(N, T, m)

	ships := strings.Split(S, ",")

	for _, ship := range ships {
		rowStart, colStart, rowEnd, colEnd := getShipArea(ship, m)
		hits := 0
		shipParts := 0

		for i := rowStart; i <= rowEnd; i++ {
			for j := colStart; j <= colEnd; j++ {
				shipParts++
				if hitMap[i][j] {
					hits++
				}
			}
		}

		switch {
		case hits == shipParts:
			sunk++
		case hits != 0:
			survived++
		}
	}

	return fmt.Sprintf("%v,%v", sunk, survived)

}

func getShipArea(ship string, m map[string]int) (int, int, int, int) {
	shipDesc := strings.Split(ship, " ") // ["1B", "2C"]

	shipHeadRow := 0
	shipHeadCol := 0

	shipTailRow := 0
	shipTailCol := 0
	for i := 0; i < 2; i++ {
		l := len(shipDesc[i])
		if i == 0 {
			shipHeadRow = m[shipDesc[i][:l-1]]
			shipHeadCol = m[shipDesc[i][l-1:]]
		} else {
			shipTailRow = m[shipDesc[i][:l-1]]
			shipTailCol = m[shipDesc[i][l-1:]]
		}
	}

	// when definition in S not in order , eq: 2C,1B

	if shipTailRow < shipHeadRow {
		shipHeadRow, shipTailRow = shipTailRow, shipHeadRow
	}

	if shipTailCol < shipHeadCol {
		shipTailCol, shipHeadCol = shipHeadCol, shipTailCol
	}
	return shipHeadRow, shipHeadCol, shipTailRow, shipTailCol
}

func getHitMap(N int, T string, m map[string]int) [][]bool {
	hitMap := make([][]bool, N)
	for i := range hitMap {
		hitMap[i] = make([]bool, N)
	}

	hits := strings.Split(T, " ")

	for i := range hits {
		l := len(hits[i]) //2 or 3  eq: 1A 26Z
		col := m[hits[i][l-1:]]
		row := m[hits[i][:l-1]]

		hitMap[row][col] = true
	}
	return hitMap
}
