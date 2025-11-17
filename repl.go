package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words

	return words
}

func firstWordREPL() error {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		if scanner.Err() != nil {
			err := fmt.Errorf("error encountered: %v", scanner.Err())
			return err
		}
		if scanner.Text() == "" {
			return fmt.Errorf("no text in the scanner")
		}
		words := cleanInput(scanner.Text())
		line := fmt.Sprintf("Your command was: %s", words[0])
		fmt.Println(line)

	}

}
