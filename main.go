/*
 *                 Copyright WavyCat 2024.
 * Distributed under the Boost Software License, Version 1.0.
 *        (See accompanying file LICENSE or copy at
 *          https://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const version = "1.3.0"

func main() {
	// Open the file with the code
	args := os.Args

	if len(args) < 2 {
		fmt.Println("No file provided")
		os.Exit(1)
	}

	if args[1] == "version" {
		fmt.Printf("GoBrainfuck v%s\nBoost Software License 1.0\n"+
			"https://github.com/wavy-cat/GoBrainfuck/releases/tag/v%s\n",
			version, version)
		os.Exit(0)
	}

	data, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	data = []byte(strings.ReplaceAll(strings.ReplaceAll(string(data), " ", ""), "\n", ""))

	// Allocate memory and create necessary variables for the program
	reader := bufio.NewReader(os.Stdin)
	memory := make([]byte, 30000)
	loopStack := make([]int, 0, 16)
	var cursor uint
	var pos int
	var printUsed bool

	if len(args) > 2 {
		num, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			fmt.Println("The standard memory size of 30,000 cells will be used")
		} else if num < 1 {
			fmt.Println("Memory size is too small. Please enter a number greater than 1.")
			os.Exit(1)
		} else {
			memory = make([]byte, num)
		}
	}

	// Executing the code
	for pos < len(data) {
		switch data[pos] {
		case '>':
			if cursor == uint(len(memory)-1) {
				fmt.Println("Cursor overflow. Details:")
				fmt.Printf("Cursor: %d. Target: %d. Memory size: %d. Code position (without LF and spaces): %d\n",
					cursor, cursor+1, len(memory), pos)
				os.Exit(1)
			}

			cursor++
		case '<':
			if cursor == 0 {
				fmt.Println("Cursor underflow. Details:")
				fmt.Printf("Cursor: %d. Target: -1. Memory size: %d. Code position (without LF and spaces): %d\n",
					cursor, len(memory), pos)
				os.Exit(1)
			}

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
			input, err := reader.ReadString('\n')
			input = strings.TrimSuffix(input, "\n")

			if err != nil {
				fmt.Println("Error reading:", err)
				os.Exit(1)
			}

			if len(input) > 2 && input[0:2] == "//" {
				i, err := strconv.Atoi(input[2:])
				if err != nil {
					fmt.Println("Error parsing number:", err)
					os.Exit(1)
				}
				memory[cursor] = byte(i)
			} else {
				if len(input) == 0 {
					memory[cursor] = 0
				} else {
					memory[cursor] = []byte(input)[0]
				}
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
