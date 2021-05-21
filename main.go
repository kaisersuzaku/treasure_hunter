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

var (
	wall        string
	path        string
	marker      string
	initPosMark string

	minStepA int
	minStepB int
	minStepC int
)

func main() {

	if len(os.Args) == 1 {
		log.Println(`Missing args. Example Use with args in order (wall path marker initPosMark): find_treasure "#" "." "$" "X"`)
		return
	}
	usageExample := `Usage example with args in order (wall path marker initPosMark) : find_treasure "#" "." "$" "X"`
	if os.Args[1] == "h" || os.Args[1] == "-h" || os.Args[1] == "--h" {
		log.Println(usageExample)
		return
	}
	if len(os.Args) > 1 && len(os.Args) < 5 {
		log.Println("Missing constraints, marker, or init pos")
		return
	}

	// Set Constraints
	wall = os.Args[1]
	path = os.Args[2]
	marker = os.Args[3]
	initPosMark = os.Args[4]

	arr := ReadArrayFromText()
	arrOut := cloneTwoDimensionStringArray(arr)

	// Find and Mark possible treasure coordinates
	coordinates := FindTreasure(arr)
	if len(coordinates) < 1 {
		fmt.Println("No possible treasure coordinates found")
	} else {
		fmt.Println()
		fmt.Println("Marked map with possible treasure coordinates : ")
		fmt.Println()
		MarkLocationAsPossibleTreasure(marker, arrOut, coordinates)
		printSlice(arrOut)
	}
}

// FindTreasure ...
// returns list of possible treasuse coordinates in rows and columns format.
// prints array result.
func FindTreasure(arr [][]string) (coordinates []string) {

	initRow, initCol, err := GetInitialPosition(arr, initPosMark)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Starting Position (row,col): %d,%d\n", initRow, initCol)
	fmt.Println()
	var a, b, c int
	var currentPosUp, currentPosRight, currentPosDown int
	for {
		a++
		currentPosUp = initRow - a
		if !IsAStepValid(currentPosUp, initCol, arr, wall, path) {
			break
		} else {
			for {
				b++
				currentPosRight = initCol + b
				if !IsBStepValid(currentPosUp, currentPosRight, arr, wall, path) {
					b = 0
					break
				} else {
					for {
						c++
						currentPosDown = currentPosUp + c
						if !IsCStepValid(currentPosDown, currentPosRight, arr, wall, path) {
							c = 0
							break
						}
						fmt.Printf("valid A B C step : %d %d %d | location (row,col) : (%d,%d)\n", a, b, c, currentPosDown, currentPosRight)
						coordinates = append(coordinates, fmt.Sprintf("%d,%d", currentPosDown, currentPosRight))
					}
				}
			}
		}
	}
	return coordinates
}

func MarkLocationAsPossibleTreasure(marker string, arrOut [][]string, coordinates []string) {
	for i := 0; i < len(coordinates); i++ {
		coordinate := strings.Split(coordinates[i], ",")
		row, _ := strconv.Atoi(coordinate[0])
		col, _ := strconv.Atoi(coordinate[1])
		arrOut[row][col] = marker
	}
}

func IsAStepValid(row, col int, arr [][]string, wall, path string) bool {
	if row >= 0 {
		if arr[row][col] == wall {
			return false
		}
		if arr[row][col] == path {
			return true
		}
		log.Panic("Hit some unknown constraints")
	}
	return false
}

func IsBStepValid(row, col int, arr [][]string, wall, path string) bool {
	if col < len(arr[row]) {
		if arr[row][col] == wall {
			return false
		}
		if arr[row][col] == path {
			return true
		}
		log.Panic("Hit some unknown constraints")
	}
	return false
}

func IsCStepValid(row, col int, arr [][]string, wall, path string) bool {
	if row < len(arr) {
		if arr[row][col] == wall {
			return false
		}
		if arr[row][col] == path {
			return true
		}
		log.Panic("Hit some unknown constraints")
	}
	return false
}

func GetInitialPosition(s [][]string, marker string) (row, col int, e error) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == marker {
				row = i
				col = j
				return
			}
		}
	}
	return 0, 0, errors.New("Init Pos not found")
}

func ReadArrayFromText() [][]string {
	file, err := os.Open("treasure_map.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	var arr [][]string
	for scanner.Scan() {
		columnLength := len(fmt.Sprint(scanner.Text()))
		col := make([]string, columnLength)
		for i := 0; i < columnLength; i++ {
			col[i] = string([]rune(fmt.Sprint(scanner.Text()))[i])
		}
		arr = append(arr, col)
	}
	return arr
}

func printSlice(s [][]string) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			fmt.Printf("%s", s[i][j])
		}
		fmt.Println()
	}
}

func cloneTwoDimensionStringArray(input [][]string) (output [][]string) {
	duplicate := make([][]string, len(input))
	for i := range input {
		duplicate[i] = make([]string, len(input[i]))
		copy(duplicate[i], input[i])
	}
	return duplicate
}
