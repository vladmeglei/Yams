package main

import (
	"fmt"
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

//
func GetPlayerNames() (int, []string) {

	var players []string
	var noOfPlayers int

	fmt.Println("How many players are in the game?")
	fmt.Println("Players:")

	if _, err := fmt.Scan(&noOfPlayers); err != nil {
		fmt.Println("Please introduce a valid number of players")
	} else {

		fmt.Println("Number of players is set to: ", noOfPlayers)
	}

	for i := 0; i < noOfPlayers; i++ {
		fmt.Println("Please introduce the player names.")
		player, _ := reader.ReadString('\n')
		players = append(players, player)

	}
	return noOfPlayers, players
}
