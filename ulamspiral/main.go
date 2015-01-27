package main

import (
	"fmt"

	"github.com/Soreil/eulersolutions/utils"
)

type point struct {
	x int
	y int
}

func main() {

	n := 11 //always odd
	lim := n * n

	primes := utils.Sieve(lim)
	spiral := make([][]int, n)
	for i := range spiral {
		spiral[i] = make([]int, n)
		for j := range spiral[i] {
			spiral[i][j] = 0
		}
	}

	start := n / 2 // center
	//if done go right
	//increment shell count
	//go shell count times-1 up
	//go shell count times left
	//go shell count times down
	//go shell count times right
	//.....

	pos := point{start, start}
	num := 1
	spiral[pos.y][pos.x+1] = 2
	for i := 2; i <= lim; i++ {
		pos.x++
		num += 2
		for j := 0; j < num-2; j++ {
			pos.y--
			i++
			if utils.IsPrime(i, primes) {
				spiral[pos.y][pos.x] = i
			}
		}
		for j := 0; j < num-1; j++ {
			pos.x--
			i++
			if utils.IsPrime(i, primes) {
				spiral[pos.y][pos.x] = i
			}
		}
		for j := 0; j < num-1; j++ {
			pos.y++
			i++
			if utils.IsPrime(i, primes) {
				spiral[pos.y][pos.x] = i
			}
		}
		for j := 0; j < num-1; j++ {
			pos.x++
			i++
			if utils.IsPrime(i, primes) {
				spiral[pos.y][pos.x] = i
			}
		}
	}
	for _, v := range spiral {
		fmt.Printf("%3d\n", v)
	}
}
