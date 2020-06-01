package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Set up players.
	var player, bot Player
	player.SetHuman()
	bot.SetBot()
	// Initialize the game map.
	MapSectors.Init()
	MapSectors.ShowMap()
	player.InitPos()
	bot.InitPos()
	MapSectors.ShowMap()
	// Players make their moves until meet each other
	for player.Location != bot.Location {
		player.Move()
		bot.Move()
		MapSectors.ShowMap()
	}
	fmt.Println("Bye.")
}
