package main

import (
	"bufio"
	"fmt"
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

func main() {

	var noOfPlayers int
	var playerNames []string
	var name string

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
		playerNames = append(playerNames, player)

	}

	for _, v := range playerNames {
		name = CreatePlayerTable(v)
	}

	fmt.Println(name)

}
