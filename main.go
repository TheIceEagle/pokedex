package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main(){
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Println(err)
		}
		input := scanner.Text()
		
		firstWrd := cleanInput(input)
		fmt.Println("Your command was:",firstWrd[0])

	}	

}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText) // Split the string by whitespace
	var result []string
	for _, word := range words {
		cleanedWord := strings.ReplaceAll(word, ",", "") // Remove commas
		result = append(result, cleanedWord)
	}
	return result
}