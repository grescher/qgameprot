package main

import (
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
	player.InitPos()
	bot.InitPos()

}
