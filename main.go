package main

import "fmt"

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
	playerNames  []string
	player1      string
	player2      string
	player3      string
	answTable    string
	answCategory string
	answScore    int
)

func main() {

	//Create players

	//Create tables

	//Start game
	//  For - number of turns
	//		Each player take turn and rolls dice
	//			Reroll 2 times if necessary
	// 				Add score to table

	playerNames = GetPlayerNames()
	player1 = playerNames[0]
	player2 = playerNames[1]
	player3 = playerNames[2]

	type player1TableM = Table
	type player2TableM = Table
	type player3TableM = Table
	type player1TableF = Table
	type player2TableF = Table
	type player3TableF = Table

	for i := 0; i < 28; i++ {
		fmt.Println("Player " + playerNames[0] + "turn: ")
		dice1 = rollDice()
		dice2 = rollDice()
		dice3 = rollDice()
		dice4 = rollDice()
		dice5 = rollDice()
		displayDiceArt()

		fmt.Println("Do you want to reroll?")
		reroll1, err := reader.ReadString('\n')
		if reroll1 == "Yes" || reroll1 == "YES" || reroll1 == "yes" {
			rerollDice()
		}
		fmt.Println("Do you want to reroll?")
		reroll2, err := reader.ReadString('\n')
		if reroll2 == "Yes" || reroll2 == "YES" || reroll2 == "yes" {
			rerollDice()
		}
		fmt.Println("What is the final score of the roll?")
		fmt.Scan(&answScore)
		fmt.Println("In which table do you want to put the score?(M - for Mandatory Table; F - for Free Table)")
		fmt.Scan(&answTable)
		fmt.Println("In which category do you want to put the score?(Select a category from the table)")
		fmt.Println("One - Two - Three - Four - Five - Six - OneP- TwoP - ThreeK - SmallF - BigF - FullH - FourK - Yams")
		fmt.Scan(&answCategory)

		if answTable == "M" {
			AddScoreToTable(answCategory, answScore, player1TableM)
		} else if answTable == "F" {
			AddScoreToTable(answCategory, answScore, player1TableF)
		}

		fmt.Printf(player1TableF)

	}

}
