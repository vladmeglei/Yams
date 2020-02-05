package main

import (
	"bufio"
	"os"
)

// When the game starts it prompts for how many players. Then it needs the name for the players to initialize the tables.
// It prints the valid combinations and then starts the first round of the game.
// Each player takes turns to roll the dice and have 2 chances to reroll if they want (any number of dices).
// After the 3rd reroll, the player chooses in which category to put the result and the score is then added to the specified table.
// Next players turn, same thing for each player until the end of the round.

//Reguli: Tabelul fiecarui jucator are 2 coloane, una pentru scoruri obligate - care trebuie completat obligatoriu in ordinea de pe tabel
// si alta pentru libere, unde poti pune scorul in orice ordine.

//Valid results:
//Yams - Five of a kind
//Four of a kind - four of a kind
//Fullhouse - 3 & 2
//Big flush - 2,3,4,5,6
//Small flush - 1,2,3,4,5
//Three of a kind - 3 dices same number
//Two pairs
//One pair

var (
	playerNames   []string
	player1       string
	player2       string
	player3       string
	answTable     string
	answCategory  string
	answScore     int
	player1TableM Table
	player1TableF Table
	name          string
	reader        = bufio.NewReader(os.Stdin)
	noPlayers     int
)

func main() {

	//Create players

	//Create tables

	//Start game
	//  For - number of turns
	//		Each player take turn and rolls dice
	//			Reroll 2 times if necessary
	// 				Add score to table

	noPlayers, playerNames = GetPlayerNames()

	//player2 = playerNames[1]
	//player3 = playerNames[2]

	// var player2TableM Table
	// var player3TableM Table
	// var player2TableF Table
	// var player3TableF Table

	for i := 0; i < 28; i++ {

	}

}
