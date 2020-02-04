package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	diceOne = `
    _______
   |       |
   |   *   |
   |_______|
   `
	diceTwo = `
    _______
   |       |
   | *   * |
   |_______|
   `
	diceThree = `
    _______
   | *     |
   |   *   |
   |_____*_|
   `
	diceFour = `
    _______
   | *   * |
   |       |
   |_*___*_|
   `
	diceFive = `
    _______
   | *   * |
   |   *   |
   |_*___*_|
   `
	diceSix = `
    _______
   | *   * |
   | *   * |
   |_*___*_|
   `
	dice1 int
	dice2 int
	dice3 int
	dice4 int
	dice5 int
)

func main() {
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

}

func rollDice() int {
	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(5) + 1

	return val
}
func displayDiceArt() {
	fmt.Println("Dice 1:")
	diceArt(dice1)
	fmt.Println("Dice 2:")
	diceArt(dice2)
	fmt.Println("Dice 3:")
	diceArt(dice3)
	fmt.Println("Dice 4:")
	diceArt(dice4)
	fmt.Println("Dice 5:")
	diceArt(dice5)

}

func diceArt(diceVal int) {
	switch diceVal {
	case 1:
		fmt.Println(diceOne)
	case 2:
		fmt.Println(diceTwo)
	case 3:
		fmt.Println(diceThree)
	case 4:
		fmt.Println(diceFour)
	case 5:
		fmt.Println(diceFive)
	case 6:
		fmt.Println(diceSix)
	}
}

func rerollDice() {
	fmt.Println("How many dices do you want to reroll?")
	var numberOfDices int
	var diceNumber int
	fmt.Scan(&numberOfDices)

	for i := 0; i < numberOfDices; i++ {
		fmt.Println("Which dice do you want to reroll?(Use the DICE number, then press ENTER)")
		fmt.Scan(&diceNumber)
		switch diceNumber {
		case 1:
			dice1 = rollDice()
		case 2:
			dice2 = rollDice()
		case 3:
			dice3 = rollDice()
		case 4:
			dice4 = rollDice()
		case 5:
			dice5 = rollDice()
		}
	}

	displayDiceArt()

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
