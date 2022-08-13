package main

import "fmt"

func main() {
	gamer := Gamer{
		Name:  "Yudi Setiawan",
		Games: []string{},
	}
	gamer.AddGame("Tekken 7")
	gamer.AddGame("Fifa 2022")
	gamer.AddGame("The last of us")
	gamer.AddGame("The last of us part 2")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

type Gamer struct {
	Name  string
	Games []string
}
