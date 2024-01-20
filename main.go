package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NewLine(str string) []string {
	var words []string
	word := ""
	for i := 0; i < len(str); i++ {
		if i < len(str)-1 {
			if str[i] == '\\' && str[i+1] == 'n' {
				words = append(words, word)
				word = ""
				i++
			} else {
				word += string(str[i])
			}

		} else {
			word += string(str[i])
			words = append(words, word)
		}
	}
	return words
}

func ConvAsciiArt(str string, char string) {

	var fileName string

	if len(os.Args) == 3 {
		banner := strings.ToLower(os.Args[2])
		switch banner {
		case "shadow":
			fileName = "shadow.txt"
		case "thinkertoy":
			fileName = "thinkertoy.txt"
		case "standard":
			fileName = "standard.txt"
		default:
			fmt.Println("Error: NO BANNER AVAILABLE!\nUsage: go run . [STRING] [BANNER]\nEx: go run . something standard")
			os.Exit(0)
		}

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		lineNumber := 1
		runes := []rune(str)
		for y := 0; y < 8; y++ {
			for i := 0; i < len(runes); i++ {
				lineNumber = 0
				valLetter := int(runes[i])
				linePrint := (valLetter-33)*9 + 11 + y

				for scanner.Scan() {
					lineNumber++
					if lineNumber == linePrint {
						line := scanner.Text()
						fmt.Print(line)
					}
				}

				file.Seek(0, 0)
				scanner = bufio.NewScanner(file)
			}
			fmt.Println()
		}
	} else {
		found := false
		wordCheck := ""
		count := 0
		wordCount := len(char)
		file, err := os.Open("standard.txt")
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		lineNumber := 1
		runes := []rune(str)
		for y := 0; y < 8; y++ {
			for i := 0; i < len(runes); i++ {
				lineNumber = 0
				valLetter := int(runes[i])
				linePrint := (valLetter-33)*9 + 11 + y
				for scanner.Scan() {
					lineNumber++
					if lineNumber == linePrint {
						line := scanner.Text()
						if char == "" {
							// fmt.Print(colorName + line + defColor)
							fmt.Print(line)
						} else {
							for x := i; x < len(char)+i && x < len(runes); x++ {
								wordCheck = wordCheck + string(str[x])
							}
							if wordCheck == char {
								found = true
							}
							wordCheck = ""
							if found {
								// fmt.Print(colorName + line + defColor)
								fmt.Print(line)
								count++
								if count == wordCount {
									count = 0
									found = false
								}
							} else {
								fmt.Print(line)
							}
						}
					}
				}
				file.Seek(0, 0)
				scanner = bufio.NewScanner(file)
			}
			fmt.Println()
		}
	}

}

func main() {
	if len(os.Args) <= 3 && len(os.Args) > 1 {
		var str, char string

		str = os.Args[1]

		words := NewLine(str)
		for i := 0; i < len(words); i++ {
			if words[i] != "" {
				ConvAsciiArt(words[i], char)
			} else if words[i] == "" {
				fmt.Println()
			}
		}
	} else {
		fmt.Println("PLEASE PROVIDE ARGUMENTS\nUsage: go run . [STRING] [BANNER]\nEx: go run . something standard")
		os.Exit(0)
	}

}
