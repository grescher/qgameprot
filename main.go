package main

func main() {
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
	// Players make their moves while there are free sectors on the map.
	for MapSectors.IsAvailable() {
		player.Move()
		bot.Move()
		MapSectors.ShowMap()
	}
	ShowResults(&player, &bot)
}
