package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	noOfPlayers int
	name        string
	reader      = bufio.NewReader(os.Stdin)
)

type Table struct {
	One    int
	Two    int
	Three  int
	Four   int
	Five   int
	Six    int
	Bonus  int
	OneP   int
	TwoP   int
	ThreeK int
	SmallF int
	BigF   int
	FullH  int
	FourK  int
	Yams   int
}

func GetPlayerNames() []string {

	var players []string

	fmt.Println("How many players are in the game?")
	fmt.Println("Players:")

	if _, err := fmt.Scan(&noOfPlayers); err != nil {
		fmt.Println("Please introduce a valid number of players")
	} else {

		fmt.Println("Number of players is set to: ", noOfPlayers)
	}

	for i := 0; i < noOfPlayers; i++ {
		player, _ := reader.ReadString('\n')
		players = append(players, player)

	}
	return players
}
