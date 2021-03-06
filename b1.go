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

	fmt.Printf("Dice: %v\n", dice)

	c := make(chan *something, 0)
	go findBalance(c, dice)

	seen := make(map[int]bool)

	for {
		b := <-c
		if b == nil {
			break
		}
		sort.Ints(b.left)
		cksum := sumck(b.left)
		if _, found := seen[cksum]; !found {
			fmt.Printf("%v == %v\n", b.left, b.right)

			seen[cksum] = true

			sort.Ints(b.right)
			seen[sumck(b.right)] = true // don't output right == left
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

func findBalance(ch chan *something, dice []int) {
	var left, right []int
	realbalance(ch, dice, left, right)
	close(ch)
}

func realbalance(ch chan *something, dice, left, right []int) {
	if len(dice) == 0 {
		sumleft, sumright := 0, 0
		for i := range left {
			sumleft += left[i]
			sumright += right[i]
		}
		if sumleft == sumright {
			ch <- &something{left, right}
		}
		return
	}

	nextdice := make([]int, len(dice)-1)
	var x, y int
	if len(left) == len(right) {
		x = 1
	} else {
		y = 1
	}
	nextleft := make([]int, len(left)+x)
	nextright := make([]int, len(right)+y)
	copy(nextleft, left)
	copy(nextright, right)

	last := 0
	for i := 0; i < len(dice); i++ {
		if dice[i] == last {
			continue
		}
		last = dice[i]
		k := 0
		var thru int
		for j := 0; j < len(dice); j++ {
			if j == i {
				thru = dice[j]
				continue
			}
			nextdice[k] = dice[j]
			k++
		}
		if len(left) == len(right) {
			nextleft[len(left)] = thru
		} else {
			nextright[len(right)] = thru
		}
		realbalance(ch, nextdice, nextleft, nextright)
	}
}
