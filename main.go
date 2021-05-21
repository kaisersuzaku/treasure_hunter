package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	printSlice(arr)
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
