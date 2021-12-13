# Balance a pile of dice

Found this on Twiter, with no other explanation:

![problem statement](20210913_081110.jpg?raw=true)

## Analysis

Another puzzle that just assumes too much.
In this case it assumes (I think) that the pips on the dice have some weight,
so you balance based on pip-count.

Physically, this is silly.
People go out of their way to make dice that aren't biased in the slightest,
even though pip counts aren't the same on any 2 faces.
Using physical intuition, the answer is any 6 on one side, the other 6 on the other side.

But that's not a very interesting problem.
I'll assume that we want the same number of pips on both sides of the balance.

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
$ ./b2  2 3 3 3 4 4 4 4 4 5 6 6
Dice: [2 3 3 3 4 4 4 4 4 5 6 6], 48
[2 3 3 3 4 4 5] == [4 4 4 6 6]
[2 3 3 4 4 4 4] == [3 4 5 6 6]
[2 3 3 4 6 6] == [3 4 4 4 4 5]
[2 3 4 4 5 6] == [3 3 4 4 4 6]
[2 4 4 4 4 6] == [3 3 3 4 5 6]
```
