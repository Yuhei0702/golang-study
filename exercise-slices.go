package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	row := make([][]uint8, dy)
	for i := range row {
		row[i] = make([]uint8, dx)
		for j := 0; j < dy; j++ {
			row[i][j] = uint8(i*j)
		}
	}

	return row
}

func main() {
	pic.Show(Pic)
}
