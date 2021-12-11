package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type something struct {
	left  []int
	right []int
}

func main() {
	var dice []int
	for _, str := range os.Args[1:] {
		n, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		dice = append(dice, n)
	}
	sort.Ints(dice)

	sum := 0
	for _, val := range dice {
		sum += val
	}

	fmt.Printf("Dice: %v, %d\n", dice, sum)
	if (sum % 2) == 1 {
		fmt.Printf("no balance possible\n")
		return
	}

	c := make(chan *something, 0)
	go findBalance(c, dice, sum/2)

	seen := make(map[int]bool)

	for {
		b := <-c
		if b == nil {
			break
		}
		sort.Ints(b.left)
		sort.Ints(b.right)
		cksumL := sumck(b.left)
		cksumR := sumck(b.right)
		if !(seen[cksumL] && seen[cksumR]) {
			var sumLeft, sumRight int
			for i := range b.left {
				sumLeft += b.left[i]
			}
			for i := range b.right {
				sumRight += b.right[i]
			}
			fmt.Printf("%v == %v\t%d == %d\n", b.left, b.right, sumLeft, sumRight)
			seen[cksumL] = true
			seen[cksumR] = true
		}
	}
}

// sumck finds base-7 value of a number which has digits of the array values.
// Array values are 1 - 6
func sumck(a []int) int {
	cksum := 0
	place := 1
	for _, val := range a {
		cksum += val * place
		place *= 7
	}
	return cksum
}

func findBalance(ch chan *something, dice []int, half int) {
	var left, right []int
	realbalance(ch, dice, left, right, half)
	close(ch)
}

func realbalance(ch chan *something, dice, left, right []int, half int) {
	sumleft, sumright := 0, 0
	for i := range left {
		sumleft += left[i]
	}
	for i := range right {
		sumright += right[i]
	}

	if len(dice) == 0 {
		if sumleft == sumright {
			ch <- &something{left, right}
		}
		return
	}

	nextdice := make([]int, len(dice)-1)
	nextleft := make([]int, len(left)+1)
	nextright := make([]int, len(right)+1)
	copy(nextleft, left)
	copy(nextright, right)
	lengthRight := len(right)
	lengthLeft := len(left)

	last := 0
	for i := 0; i < len(dice); i++ {
		if dice[i] == last {
			continue
		}
		last = dice[i]
		var thru int
		for j, k := 0, 0; j < len(dice); j++ {
			if j == i {
				thru = dice[i]
				continue
			}
			nextdice[k] = dice[j]
			k++
		}

		if sumleft+thru <= half {
			nextleft[lengthLeft] = thru
			realbalance(ch, nextdice, nextleft, nextright[:lengthRight], half)
		}

		if sumright+thru <= half {
			nextright[lengthRight] = thru
			realbalance(ch, nextdice, nextleft[:lengthLeft], nextright, half)
		}
	}
}
