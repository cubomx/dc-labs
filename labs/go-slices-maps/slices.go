package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var multi = make([][]uint8, dy*3)
	for i := 0; i < dy; i++ {
		multi[i] = make([]uint8, dx*3)
		for j := 0; j < dx; j++ {
			multi[i][j] = uint8((i+j)/2)
		}
	}
	return multi[:dy]
}

func main() {
	pic.Show(Pic)
}
