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

	c := make(chan *something, 20)
	go findBalance(c, dice, sum/2)

	seen := make(map[int]bool)

	for {
		b := <-c
		if b == nil {
			break
		}
		sort.Ints(b.left)
		cksumL := sumup(b.left)
		if !seen[cksumL] {
			sort.Ints(b.right)
			cksumR := sumup(b.right)
			if !seen[cksumR] {
				fmt.Fprintf(os.Stdout, "%v == %v\n", b.left, b.right)
				seen[cksumL] = true
				seen[cksumR] = true
			}
		}
	}
}

// sumup calculates a single value for values []int
// as if the slice values were digits of a base-7
// number representation.
func sumup(values []int) int {
	sum := 0
	for i := range values {
		sum = sum*7 + values[i]
	}
	return sum
}

func findBalance(ch chan *something, dice []int, half int) {
	var left, right []int
	realbalance(ch, dice, left, right, 0, 0, half)
	close(ch)
}

func realbalance(
	ch chan *something,
	dice, left, right []int,
	sumleft, sumright, half int) {

	if sumleft == half && sumleft == sumright {
		nleft := make([]int, len(left))
		copy(nleft, left)
		nright := make([]int, len(right))
		copy(nright, right)
		ch <- &something{left: nleft, right: nright}
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
		nextdice = nextdice[:0]
		nextdice = append(nextdice, dice[:i]...)
		nextdice = append(nextdice, dice[i+1:]...)
		thru := dice[i]

		if sumleft+thru <= half {
			nextleft[lengthLeft] = thru
			realbalance(ch, nextdice, nextleft, nextright[:lengthRight], sumleft+thru, sumright, half)
		}

		if sumright+thru <= half {
			nextright[lengthRight] = thru
			realbalance(ch, nextdice, nextleft[:lengthLeft], nextright, sumleft, sumright+thru, half)
		}
	}
}
