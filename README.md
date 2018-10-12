# Go Poker
[![Go Report Card](https://goreportcard.com/badge/github.com/syllabix/go-poker)](https://goreportcard.com/report/github.com/syllabix/go-poker)

A project for analyzing and projecting poker hands

## Overview
This project is designed more as an exploration in optimizing Go code while still maintaining something readable and scalable.

It contains features to analyze poker hands, as well determine the best possible hand in a 5 card draw style game knowing the cards at the top of the deck.

The project expects cards to represented in 2 characters code format symbolizing Rank and Suit:

Ranks:
2-9
T = 10,
J = Jack
Q = Queen
K = King
A = Ace

Suits:
S = Spades
C = Clubs
D = Diamonds
H = Hearts

Example Hand:
`2H 2S 3H 3S 3C`

### Running the project
The following instructions assume Go is installed and a Go workspace is set up.

1. Clone the repository
2. Navigate to the root of the project and run `go run main.go`
3. Follow the interactive terminal prompts

```sh
Example Input:
2H 2S 3H 3S 3C 2D 9C 3D 6C TH
^------------^ ^------------^
     Hand           Deck
```


#### Benchmarks
```sh
goos: darwin
goarch: amd64
pkg: github.com/syllabix/psychic-poker-player/poker
BenchmarkGetRank/Straight_Flush-8   392 ns/op    3 allocs/op
BenchmarkGetRank/Four_of_a_Kind-8   314 ns/op    3 allocs/op
BenchmarkGetRank/Full_House-8       324 ns/op    3 allocs/op
BenchmarkGetRank/Flush-8            420 ns/op    3 allocs/op
BenchmarkGetRank/Straight-8         401 ns/op    3 allocs/op
BenchmarkGetRank/Three_of_a_Kind-8  335 ns/op    3 allocs/op
BenchmarkGetRank/Two_Pair-8         343 ns/op    3 allocs/op
BenchmarkGetRank/One_Pair-8         356 ns/op    3 allocs/op
BenchmarkGetRank/High_Card-8        404 ns/op    3 allocs/op

```