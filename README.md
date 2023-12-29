# Balance a pile of dice

Found this on Twitter, with no other explanation:

![problem statement](20210913_081110.jpg?raw=true)

I'll state it in text so others can find it:

```
Put these dice on the scale so that it balances.
```

It's just a puzzle, not a programming job interview question.

But it is a programming job interview question!

---

### Daily Coding Problem: Problem #1105 [Medium]

This problem was asked by Facebook.

Given a multiset of integers,
return whether it can be partitioned into two subsets whose sums are the same.

For example,
given the multiset {15, 5, 20, 10, 35, 15, 10}, it would return true,
since we can split it up into {15, 5, 10, 15, 10} and {20, 35},
which both add up to 55.

Given the multiset {15, 5, 20, 10, 35}, it would return false,
since we can't split it up into two subsets that add up to the same sum.

---

This is exactly the same problem, phrased differently.
For the record, the example sets of numbers partition like this:

```
$ ./b2 15 5 20 10 35 15 10
Dice: [5 10 10 15 15 20 35], 110
[5 10 10 15 15] == [20 35]
[5 15 15 20] == [10 10 35]
[5 15 35] == [10 10 15 20]
```

I'm a little skeptical that my "checksum" works on numbers other than 1-6,
but the above example seems to work.

### Daily Coding Problem: Problem #1591 [Hard]

This problem was asked by Google.

Given a list of integers S and a target number k,
write a function that returns a subset of S that adds up to k.
If such a subset cannot be made, then return null.

Integers can appear more than once in the list.
You may assume all numbers in the list are positive.

For example, given S = [12, 1, 61, 5, 9, 2] and k = 24,
return [12, 9, 2, 1] since it sums up to 24.

---

Another subset sum variant as a Daily Coding Problem,
this time a "[Hard]".
As formulated by Google,
this seems like easier coding than Facebook's variant,
which is only "[Medium]".
There's no accounting for taste.

## Analysis

Another puzzle that just assumes too much.
In this case it assumes (I think) that the pips on the dice have some weight,
so you balance based on pip-count.

Physically, this is silly.
People go out of their way to make dice that aren't biased in the slightest,
even though pip counts aren't the same on any 2 faces.
Using physical intuition, the answer is any 6 on one side, the other 6 on the other side.

A solution that simple isn't a very interesting problem.
I'll assume that we want the same number of pips on both sides of the balance.

This turns out to be a formulation of the [Subset-Sum Problem](https://en.wikipedia.org/wiki/Subset_sum_problem),
which is said to be NP-complete.
I don't feel too bad about my solution.

```sh
$ go build b1.go
$ ./b1 2 3 3 3 4 4 4 4 4 5 6 6 
Dice: [2 3 3 3 4 4 4 4 4 5 6 6]
[2 3 4 4 5 6] == [3 3 4 4 4 6]
[2 3 3 4 6 6] == [3 4 4 4 4 5]
[2 4 4 4 4 6] == [3 3 3 4 5 6]
```

Looks like I also assumed that equal numbers of dice
have to be on both sides of the balance.

```sh
$ go build b2.go
$ time ./b2  2 3 3 3 4 4 4 4 4 5 6 6

Dice: [2 3 3 3 4 4 4 4 4 5 6 6], 48
[2 3 3 3 4 4 5] == [4 4 4 6 6]
[2 3 3 4 4 4 4] == [3 4 5 6 6]
[2 3 3 4 6 6] == [3 4 4 4 4 5]
[2 3 4 4 5 6] == [3 3 4 4 4 6]
[2 4 4 4 4 6] == [3 3 3 4 5 6]
./b2 2 3 3 3 4 4 4 4 4 5 6 6  67.94s user 4.21s system 122% cpu 58.710 total
```

## Design

I used Go to write the programs.
Go has built-in concurrency primitives,
and it's quite easy to do concurrent programs.
I wrote a recursive backtracking algorithm to generate
left and right sides of the balance.
That recursive algorithm ran in its own goroutine,
passing balanced sets of dice back to the main goroutine via a Go channel.
The main goroutine reads balanced sets of dice from the channel,
and decides whether it has already encountered any given balanced set.

Since this is explicitly done in terms of ordinary, cubical, D6 dice,
I sorted the arrays of left and right side dice numerically,
then treated the arrays as digits of a base-7 number.
The slice-of-ints `[2 4 4 4 4 6]` gets treated as a base-7 number
that would have a text representation of 644,442<sub>7</sub>,
and a base-10 text representation of 114844.
I used these values as identifiers of particular balanced sets of dice,
and could check if the main goroutine had encountered a particular "balance" already.

This concurrency works well.
The recursive, backtracking code does not have to have the clutter of doing the
work of deciding if a balance has been seen before, and doing the output.
The main goroutine doesn't do anything other than receive balanced sets of integers,
then decide if it should write them to output or not.
