package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Types */

// Player type.
type Player struct {
	Name     string
	Location Coords // Tracks in which sector the player currently is.
	Sectors  int    // Counts for the number of sectors captured.
	IsBot    bool   // True if a player is bot.
}

/* Variables for the local use */

// initHumPos keeps the initial sector value for the human player.
// Used in the bot's first sector initialization.
var initHumPos int

/* Package methods */

// SetBot marks the player as bot.
func (p *Player) SetBot() {
	p.Name = "Bot Player"
	p.IsBot = true
	rand.Seed(time.Now().UnixNano())
}

// SetHuman marks the player as human.
func (p *Player) SetHuman() {
	p.Name = "Human Player"
	p.IsBot = false
}

// InitPos sets up the initial player's sector.
func (p *Player) InitPos() {
	var choice int
	if !p.IsBot {
		// The human player initial sector.
		fmt.Println("Choose sector to start:")
		choice := selectSector(MapMax)
		initHumPos, p.Location = choice, ToCoords(choice)
	} else {
		// The bot player initial sector.
		choice = rand.Intn(MapMax) + 1
		for choice == initHumPos {
			// Reassign location if it's equal to the human player location.
			choice = rand.Intn(MapMax) + 1
		}
		p.Location = ToCoords(choice)
		fmt.Printf("The opponent's clan has picked sector %d\n", choice)
	}
	p.CaptureSector(&MapSectors)
}

// Move sets up the next player's position on the map.
func (p *Player) Move() {
	var newLocation Coords
	if !p.IsBot {
		// The human player initial sector.
		fmt.Println("Input the sector number to move:")
		// do...while loop analogue
		for {
			newLocation = ToCoords(selectSector(MapMax))
			err := checkNextMove(p.Location, newLocation)
			if err != nil {
				fmt.Println(err)
				fmt.Print("Select the sector next to your current location: ")
				continue
			}
			setNewLocation(p, newLocation)
			break
		}
	} else {
		for {
			newLocation = ToCoords(rand.Intn(MapMax) + 1)
			err := checkNextMove(p.Location, newLocation)
			if err != nil {
				continue
			}
			setNewLocation(p, newLocation)
			break
		}
	}
}

// CaptureSector sets the status of the map sector for the respective player.
func (p *Player) CaptureSector(m *Map) {
	if p.IsBot {
		if m[p.Location.Row][p.Location.Col].Status == SectEmp {
			m[p.Location.Row][p.Location.Col].Status += SectBot
			p.Sectors++
		}
	} else {
		if m[p.Location.Row][p.Location.Col].Status == SectEmp {
			m[p.Location.Row][p.Location.Col].Status += SectHum
			p.Sectors++
		}
	}
}
