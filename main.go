/*
	Cole Zandbergen
	CS424 Fall 2021
	Programming Assignment 1

	*********************
	This program will  read in
	Baseball player objects from
	a text file. This code will
	assume there are no errors in
	the file
	*********************
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/****
	This is where the Player object is defined
****/

type Player struct {
	//All statistics that are stored in the player object
	firstname          string
	lastname           string
	plateappearances   float64
	atbats             float64
	singles            float64
	doubles            float64
	triples            float64
	homeruns           float64
	walks              float64
	hitbypitch         float64
	battingaverage     float64
	sluggingpercentage float64
	obp                float64 //this stands for on-base percentage
	ops                float64 //this stands for on-base plus slugging
}

// function to initialize a player object
// the string input is intended to be a line from the data file
func (p *Player) Initialize(s string) {
	list := strings.Split(s, " ") //split the string into different words, then assign each of
	p.firstname = list[0]         //the player's attributes based on the list of words
	p.lastname = list[1]          //this relies on the inputs being in the correct order
	p.plateappearances = convert(list[2])
	p.atbats = convert(list[3])
	p.singles = convert(list[4])
	p.doubles = convert(list[5])
	p.triples = convert(list[6])
	p.homeruns = convert(list[7])
	p.walks = convert(list[8])
	p.hitbypitch = convert(list[9])

	//Set all the variables that are listed in the file, then call the functions to set the stats
	//that need to be computed
	p.setBA()
	p.setSP()
	p.setOBP()
	p.setOPS()
}

func (p Player) print() {
	name := p.lastname + ", " + p.firstname
	fmt.Printf("\n%20s%-5s%1.3f%1.3f%1.3f%1.3f", name, ":", p.battingaverage, p.sluggingpercentage, p.obp, p.ops)
}

//this function converts a string to an int
//This function's existence makes it easier for me to initialize each property of the player object
func convert(s string) float64 {
	var x float64
	x, _ = strconv.ParseFloat(s, 64)
	return x
}

//function to calculate and set the player's batting average
func (p *Player) setBA() {
	a := p.singles + p.doubles + p.triples + p.homeruns
	fmt.Printf("\nTotal hits for %s: %f", p.lastname, a)
	fmt.Printf("\nTotal atbats for %s: %f", p.lastname, p.atbats)
	p.battingaverage = a / p.atbats
	fmt.Printf("\nBatting average for %s: %1.3f", p.lastname, p.battingaverage)
}

//function to calculate and set a player's slugging percentage
func (p *Player) setSP() {
	s := p.singles
	d := p.doubles * 2
	t := p.triples * 3
	h := p.homeruns * 4
	p.sluggingpercentage = (s + d + t + h) / p.atbats
}

//function to calculate and set a player's on base percentage
func (p *Player) setOBP() {
	s := p.singles
	d := p.doubles
	t := p.triples
	h := p.homeruns
	w := p.walks
	hbp := p.hitbypitch
	total := s + d + t + h + w + hbp
	p.obp = total / p.plateappearances
}

//function to calculate and set a player's on-base plus slugging percentage
func (p *Player) setOPS() {
	//this one's easy, just set ops to be sluggingpercentage + obp
	p.ops = p.sluggingpercentage + p.obp
}

/**
	Main function
**/
func main() {

	//Declare and initialize variables here
	var input *bufio.Scanner
	var fileReader *bufio.Scanner
	var file *os.File
	var filename string
	var e error
	var PlayerList []Player

	//Print welcome message and prompt for input
	fmt.Println("Welcome, this program will calculate and display baseball player statistics from a formatted input file.\n First, enter the name of the file")
	//Now, initialize the scanner, and use it to retrieve input from the user
	input = bufio.NewScanner(os.Stdin)
	input.Scan()
	//Now that we have the input, we will use it to try and open a file
	filename = input.Text()
	file, e = os.Open(filename)
	if e != nil { //If the file cannot be opened, then we will notify the user and exit the program
		fmt.Println("There was an error opening the file.")
		fmt.Println("Exiting the program")
		os.Exit(1)
	} else {
		//If there are no errors, it's time to get started
		//First, create a scanner for the file
		fileReader = bufio.NewScanner(file)
		for fileReader.Scan() { //Read each line from the file until we reach the end of the file
			//create a temporary player object, then append it to the Player list after it is initialized
			t := new(Player)
			t.Initialize(fileReader.Text())
			PlayerList = append(PlayerList, *t)
		}

		//Now that the list is complete, print out each player's stats
		fmt.Printf("Baseball team report: --- %d players were found in the file", len(PlayerList))
		fmt.Printf("\n%-20s%-5s%15s%15s%15s%15s", "     PLAYER NAME", ":", "AVERAGE", "SLUGGING", "ONBASE%", "OPS")

		for i := 0; i < len(PlayerList); i++ {
			PlayerList[i].print()
		}
	}
}
