package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	noOfPlayers int
	name        string
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

// sss
func CreatePlayerTable(playerName string) string {
	var playerTable string
	playerTable = playerName
	return playerTable
}

func GetPlayerNames() []string {

	var players []string

	reader := bufio.NewReader(os.Stdin)
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

func AddScoreToTable(category string, score int, table Table) {

	switch category {
	case "One":
		table.One = score
		fmt.Println("Added score ", score, " to category One")
	case "Two":
		table.Two = score
		fmt.Println("Added score ", score, " to category Two")
	case "Three":
		table.Three = score
		fmt.Println("Added score ", score, " to category Three")
	case "Four":
		table.Four = score
		fmt.Println("Added score ", score, " to category Four")
	case "Five":
		table.Five = score
		fmt.Println("Added score ", score, " to category Five")
	case "Six":
		table.Six = score
		fmt.Println("Added score ", score, " to category Six")
	case "OneP":
		table.OneP = score + 10
		fmt.Println("Added score ", score, " to category One Pair")
	case "TwoP":
		table.TwoP = score + 20
		fmt.Println("Added score ", score, " to category Two Pairs")
	case "ThreeK":
		table.ThreeK = score + 30
		fmt.Println("Added score ", score, " to category Three of a Kind")
	case "SmallF":
		table.SmallF = score
		fmt.Println("Added score ", score, " to category Small Flush")
	case "BigF", "Bigflush", "Big flush":
		table.BigF = score
		fmt.Println("Added score ", score, " to category Big Flush")
	case "FullH":
		table.FullH = score + 40
		fmt.Println("Added score ", score, " to category FullHouse")
	case "FourK", "Fourofakind", "Four of a kind":
		table.FourK = score + 50
		fmt.Println("Added score ", score, " to category Four of a Kind")
	case "Yams":
		table.Yams = score + 60
		fmt.Println("Added score ", score, " to category Yams")

	}

}
