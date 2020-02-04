package main

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

func AddScoreToTable(category string, score int, table Table) {

	switch category {
	case "One":
		table.One = score
	case "Two":
		table.Two = score
	case "Three":
		table.Three = score
	case "Four":
		table.Four = score
	case "Five":
		table.Five = score
	case "Six":
		table.Six = score
	case "OneP":
		table.OneP = score + 10
	case "TwoP":
		table.TwoP = score + 20
	case "ThreeK":
		table.ThreeK = score + 30
	case "SmallF":
		table.SmallF = score
	case "BigF", "Bigflush", "Big flush":
		table.BigF = score
	case "FullH":
		table.FullH = score + 40
	case "FourK", "Fourofakind", "Four of a kind":
		table.FourK = score + 50
	case "Yams":
		table.Yams = score + 60

	}

}
