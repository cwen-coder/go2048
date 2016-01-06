package main

import (
	"fmt"
)

type G2048 [4][4]int

func (t *G2048) MirrorV() {
	temp := new(G2048)
	for i, line := range t {
		for j, num := range line {
			temp[len(t)-i-1][j] = num
		}
	}
	*t = *temp
}

func (t *G20482048) Right90() {
	temp := new(G2048)
	for i, line := range t {
		for j, num := range line {
			temp[j][len(t)-i-1] = num
		}
	}
	*t = *temp
}

func (t *G2048) Left90() {
	temp := new(G2048)
	for i, line := range t {
		for j, num := range line {
			temp[len(t)-j-1][i] = num
		}
	}
	*t = *temp
}

func (t *G2048) Right180() {
	temp := new(G2048)
	for i, line := range t {
		for j, num := range line {
			temp[len(t)-i-1][len(t)-j-1] = num
		}
	}
	*t = *temp
}

func (t *G2048) Print() {
	for _, line := range t {
		for _, num := range line {
			fmt.Printf("%2d ", num)
		}
		fmt.Println()
	}
	tn := G2048{{1, 2, 3, 4}, {5, 8}, {9, 10, 11}, {13, 14, 16}}
	*t = tn
}

func main() {
	fmt.Println("origin")
	t := G2048{{1, 2, 3, 4}, {5, 8}, {9, 10, 11}, {13, 14, 16}}
	fmt.Println("Mirror")
	t.MirrorV()
	t.Print()
	fmt.Println("Left90")
	t.Left90()
	t.Print()
	fmt.Println("Right90")
	t.Right90()
	t.Print()
	fmt.Println("Right180")
	t.Right180()
	t.Print()
}
