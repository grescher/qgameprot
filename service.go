package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ShowResults prints how many sectors each player has captured and who's the winner.
func ShowResults(players ...*Player) {
	var winner int
	for i := range players {
		fmt.Printf("%s captured %d sectors.\n", players[i].Name, players[i].Sectors)
		if players[i].Sectors > players[winner].Sectors {
			winner = i
		}
	}
	fmt.Println("The winner is", players[winner].Name)
}

/* Functions for the local use */

func selectSector(lim int) (input int) {
	// Scan int from stdin.
	input = getInt()
	// Check whether ch is in range of map, if it's not reprompt the user.
	for input < 1 || input > lim {
		fmt.Printf("incorrect value, try with <1-%d>: ", lim)
		input = getInt()
	}
	return input
}

func getInt() (n int) {
	// Create a new bufio.Reader to read keyboard input.
	reader := bufio.NewReader(os.Stdin)
	// Read all the user input untill newline character.
	input, err := reader.ReadString('\n')
	checkErr(err)
	// Remove the whitespaces (the newline chararcter).
	input = strings.TrimSpace(input)
	n, err = strconv.Atoi(input)
	checkErr(err)
	return n
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func setNewLocation(p *Player, newLoc Coords) {
	p.Location = newLoc
	p.CaptureSector(&MapSectors)
}

// Checks if the player can make a move to a new sector.
func checkNextMove(currentLoc, newLoc Coords) (err error) {
	diff := newLoc.Col - currentLoc.Col
	if diff > 1 || diff < -1 {
		return errors.New("unable to move through a sector")
	}
	diff = newLoc.Row - currentLoc.Row
	if diff > 1 || diff < -1 {
		return errors.New("unable to move through a sector")
	}
	return nil
}
