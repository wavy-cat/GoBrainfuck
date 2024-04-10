package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file with the code
	args := os.Args

	if len(args) < 2 {
		fmt.Println("No file provided")
		os.Exit(1)
	}

	if args[1] == "version" {
		fmt.Println("GoBrainfuck v1.1.0\nBoost Software License 1.0\n" +
			"https://github.com/wavy-cat/GoBrainfuck/releases/tag/v1.1.0")
		os.Exit(0)
	}

	data, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	data = []byte(strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", ""))

	// Allocate memory and create necessary variables for the program
	memory := make([]byte, 30000)
	loopStack := make([]int, 0, 16)
	var cursor uint
	var pos int
	var printUsed bool

	if len(args) > 2 {
		num, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		} else if num < 1 {
			panic("memory size is too small")
		}
		memory = make([]byte, num)
	}

	// Executing the code
	for pos < len(data) {
		switch data[pos] {
		case '>':
			if cursor == uint(len(memory)-1) {
				fmt.Printf("Cursor: %d. Code pos (without LF and spaces): %d\n", cursor, pos)
				panic("cursor overflow")
			}

			cursor++
		case '<':
			cursor--
		case '+':
			memory[cursor]++
		case '-':
			memory[cursor]--
		case '.':
			printUsed = true
			fmt.Print(string(memory[cursor]))
		case '*':
			printUsed = true
			fmt.Print(memory[cursor])
		case ',':
			var input string
			_, err = fmt.Scanln(&input)

			if err != nil {
				panic(err)
			}

			if len(input) > 2 && input[0:2] == "//" {
				i, err := strconv.Atoi(input[2:])
				if err != nil {
					panic(err)
				}
				memory[cursor] = byte(i)
			} else {
				memory[cursor] = input[0]
			}
		case '[':
			if memory[cursor] == 0 {
				openBracketCount := 1
				for openBracketCount > 0 {
					pos++
					switch data[pos] {
					case '[':
						openBracketCount++
					case ']':
						openBracketCount--
					}
				}
			} else {
				loopStack = append(loopStack, pos)
			}
		case ']':
			if memory[cursor] != 0 {
				pos = loopStack[len(loopStack)-1]
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		}
		pos++
	}

	if printUsed {
		fmt.Print("\n")
	}
}
