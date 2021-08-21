package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type rover struct {
	x         int
	y         int
	direction byte
}

func (r *rover) String() string {
	return fmt.Sprintf("%d %d %s", r.x, r.y, string(r.direction))
}

type directionNode struct {
	left  byte
	right byte
}

// compass directions
const (
	north byte = 'N'
	east  byte = 'E'
	south byte = 'S'
	west  byte = 'W'
)

// directions
const (
	left  byte = 'L'
	right byte = 'R'
	move  byte = 'M'
)

// upper-right coordinates
var (
	xCoordinateLimit = 0
	yCoordinateLimit = 0
)

var directionNodeList map[byte]directionNode
func init() {
	directionNodeList = map[byte]directionNode{}
	directionNodeList[north] = directionNode{
		left:  west,
		right: east,
	}
	directionNodeList[east] = directionNode{
		left:  north,
		right: south,
	}
	directionNodeList[south] = directionNode{
		left:  east,
		right: west,
	}
	directionNodeList[west] = directionNode{
		left:  south,
		right: north,
	}
}

// directive rover's directive
func directive(r *rover, d byte) error {
	switch d {
	case left:
		r.direction = directionNodeList[r.direction].left
	case right:
		r.direction = directionNodeList[r.direction].right
	case move:
		switch r.direction {
		case north:
			r.y++
		case east:
			r.x++
		case south:
			r.y--
		case west:
			r.x--
		}
		if r.x > xCoordinateLimit {
			return errors.New(fmt.Sprintf("x coordinate pointed to (%d) a point outside the plateau", r.x))
		}
		if r.y > yCoordinateLimit {
			return errors.New(fmt.Sprintf("y coordinate pointed to (%d) a point outside the plateau", r.y))
		}
	default:
		return errors.New(fmt.Sprintf("unexpected directive: %s", string(d)))
	}
	return nil
}

// parseUpperRightCoordinates parses raw coordinates string
func parseUpperRightCoordinates(rawCoordinates string) (x int, y int, err error) {
	// parse raw coordinates
	upperRightCoordinates := strings.Split(rawCoordinates, " ")
	if len(upperRightCoordinates) != 2 {
		return 0, 0, errors.New("failed to parse upper-right coordinates")
	}
	// validate
	if len(upperRightCoordinates) != 2 {
		return 0, 0, errors.New("failed to parse upper-right coordinates, min length of raw coordinates should be 2")
	}
	// upper right coordinates value for x
	x, err = strconv.Atoi(upperRightCoordinates[0])
	if err != nil {
		return 0, 0, err
	}
	// upper right coordinates value for y
	y, err = strconv.Atoi(upperRightCoordinates[1])
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

// parseRoverPosition parses rover's position
func parseRoverPosition(rawPosition string) (*rover, error) {
	var err error
	position := strings.Split(rawPosition, " ")
	if len(position) != 3 {
		return nil, errors.New("failed to set first rover position, position data should contain X, Y coordinate and direction (for example: \"3 2 W\")")
	}
	r := &rover{}
	// set first rover's x coordinate
	r.x, err = strconv.Atoi(position[0])
	if err != nil {
		return nil, err
	}
	// set first rover's y coordinate
	r.y, err = strconv.Atoi(position[1])
	if err != nil {
		return nil, err
	}
	// set first rover's direction
	direction := position[2]
	if len(direction) != 1 {
		return nil, errors.New("failed to parse first rover's direction")
	}
	validateDirection := bytes.ContainsAny([]byte(direction), "NESW")
	if !validateDirection {
		return nil, errors.New(fmt.Sprintf("unexpected value for direction: %s", string([]byte(direction)[0])))
	}
	r.direction = []byte(direction)[0]
	return r, nil
}

// Run root function of program
func Run(input string) (firstRover *rover, secondRover *rover, err error) {
	// input parsing
	parseInput := strings.Split(input, "\n")
	if len(parseInput) != 5 {
		return nil, nil, errors.New("error occured while input")
	}

	xCoordinateLimit, yCoordinateLimit, err = parseUpperRightCoordinates(parseInput[0])
	if err != nil {
		return nil, nil, fmt.Errorf("error occured while parsing upper right coordinates (%s)", err.Error())
	}

	// set first rover's position
	firstRover, err = parseRoverPosition(parseInput[1])
	if err != nil {
		return nil, nil, fmt.Errorf("error occured while parsing first rover position (%s)", err.Error())
	}

	// first rover's directives
	for _, d := range []byte(parseInput[2]) {
		err = directive(firstRover, d)
		if err != nil {
			return nil, nil, fmt.Errorf("error occured while applying directive (%s)", err.Error())
		}
	}

	// set second rover's position
	secondRover, err = parseRoverPosition(parseInput[3])
	if err != nil {
		return nil, nil, fmt.Errorf("error occured while parsing first rover position (%s)", err.Error())
	}

	// second rover's directives
	for _, d := range []byte(parseInput[4]) {
		err = directive(secondRover, d)
		if err != nil {
			return nil, nil, fmt.Errorf("error occured while applying directive (%s)", err.Error())
		}
	}

	return firstRover, secondRover, nil
}

func main() {
	input := "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM"
	firstRover, secondRover, err := Run(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result of First Rover: \"%s\"\n", firstRover.String())
	fmt.Printf("Result of Second Rover: \"%s\"\n", secondRover.String())
}
