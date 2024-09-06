package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage
	//
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		lines := bytes.Split(fileContents, []byte("\n"))
		for i, line := range lines {
			n := 0
		lineLoop:
			for n < len(line) {
				switch x := string(line[n]); {
				case x == "(":
					fmt.Println("LEFT_PAREN ( null")
				case x == ")":
					fmt.Println("RIGHT_PAREN ) null")
				case x == "{":
					fmt.Println("LEFT_BRACE { null")
				case x == "}":
					fmt.Println("RIGHT_BRACE } null")

				case x == ",":
					fmt.Println("COMMA , null")

				case x == ".":
					fmt.Println("DOT . null")

				case x == "+":
					fmt.Println("PLUS + null")

				case x == "*":
					fmt.Println("STAR * null")

				case x == "-":
					fmt.Println("MINUS - null")
				case x == ";":
					fmt.Println("SEMICOLON ; null")
				case x == "=":
					if n < len(line)-1 && line[n+1] == byte('=') {
						fmt.Println("EQUAL_EQUAL == null")
						n += 1
					} else {
						fmt.Println("EQUAL = null")
					}
				case x == "!":
					if n < len(line)-1 && line[n+1] == byte('=') {
						fmt.Println("BANG_EQUAL != null")
						n += 1
					} else {
						fmt.Println("BANG ! null")
					}
				case x == "<":
					if n < len(line)-1 && line[n+1] == byte('=') {
						fmt.Println("LESS_EQUAL <= null")
						n += 1
					} else {
						fmt.Println("LESS < null")
					}
				case x == ">":
					if n < len(line)-1 && line[n+1] == byte('=') {
						fmt.Println("GREATER_EQUAL >= null")
						n += 1
					} else {
						fmt.Println("GREATER > null")
					}
				case x == "/":
					if n < len(line)-1 && line[n+1] == byte('/') {
						break lineLoop
					} else {
						fmt.Println("SLASH / null")
					}
				case x == "\"":
					sComp := false
					for j := n + 1; j < len(line); j++ {

						if line[j] == byte('"') {
							fmt.Printf("STRING \"%s\" %s\n", string(line[n+1:j]), string(line[n+1:j]))
							n = j
							sComp = true
							break
						}

					}
					if !sComp {
						fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", i+1)
						defer os.Exit(65)
						break lineLoop

					}
				case unicode.IsDigit(rune(x[0])):
					j := n
					for j < len(line) {
						if !unicode.IsDigit(rune(line[j])) && string(line[j]) != "." {
							break
						}
						j++
					}
					val := string(line[n:j])
					fmt.Printf("NUMBER %s %s\n", val, formatNumber(val))
					n = j-1
				case x == " " || x == "\t":

				default:
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", i+1, x)

					defer os.Exit(65)

				}
				n += 1
			}

		}

		fmt.Println("EOF  null")
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}

func formatNumber(s string) string {
	floatVal, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(s, ".") {
		return strings.TrimRight(s, "0")
	}
	return fmt.Sprintf("%.1f", floatVal)

}
