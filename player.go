package main

import (
	"fmt"
	"log"
	"math/rand"
)

/* Types */

// Coords type contains coorginates on the game map.
type Coords struct {
	Row int
	Col int
}

// Player type.
type Player struct {
	Location Coords // Tracks in which sector the player currently is.
	IsBot    bool   // True if a player is bot.
}

/* Variables for the local use */

// initHumPos keeps the initial sector value for the human player.
// Used in the bot's first sector initialization.
var initHumPos int

/* Package methods */

// SetBot marks the player as bot.
func (p *Player) SetBot() {
	p.IsBot = true
}

// SetHuman marks the player as human.
func (p *Player) SetHuman() {
	p.IsBot = false
}

// InitPos sets up the initial player's sector.
func (p *Player) InitPos() {
	var choice int
	if !p.IsBot {
		// The human player initial sector.
		fmt.Println("Choose sector to start:")
		choice, err := readUserInput(MapMax)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("You have picked sector %d\n", choice)
		initHumPos, p.Location = choice, toCoords(choice)
	} else {
		// The bot player initial sector.
		choice = rand.Intn(MapMax) + 1
		for choice == initHumPos {
			// Reassign location if it's equal to the human player location.
			choice = rand.Intn(MapMax) + 1
		}
		p.Location = toCoords(choice)
		fmt.Printf("The opponent's clan has picked sector %d\n", choice)
	}
	p.CaptureSector(&MapSectors)
}

// Move sets up the next player's position on the map.
func (p *Player) Move() {

	if !p.IsBot {

	} else {

	}
}

// CaptureSector sets the status of the map sector for the respective player.
func (p *Player) CaptureSector(m *Map) {
	if p.IsBot {
		m[p.Location.Row][p.Location.Col].Status += SectBot
	} else {
		m[p.Location.Row][p.Location.Col].Status += SectHum
	}
}

func readUserInput(lim int) (input int, err error) {
	// Scan int from stdin.
	_, err = fmt.Scanf("%d", &input)
	if err != nil {
		return 0, err
	}
	// Check whether ch is in range of map, if it's not reprompt the user.
	for input < 1 || input > lim {
		fmt.Printf("incorrect value, try with <1-%d>: ", lim)
		_, err = fmt.Scanf("%d", &input)
		if err != nil {
			return 0, err
		}
	}
	return input, nil
}

// Converts map sector number into map coordinates.
func toCoords(n int) (c Coords) {
	c.Row = (n - 1) / MapWidth
	c.Col = (n - 1) % MapWidth
	return c
}
