package main

import (
	"fmt"
	"movies/cmd"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"golang.org/x/exp/rand"
)

func main() {

	greetUsers()

	cmd.Execute()

}

// **Simple Ascii art greeting**
// func greetUsers() {
// 	fig := figure.NewColorFigure("WELCOME", "puffy", "white", true)
// 	fig.Print()
// 	fmt.Println("\nUse this tool to get the Movie updates you need!")
// }

// Random coloured greeting
// Function to get a random color from a predefined list
func getRandomColor() *color.Color {
	colors := []*color.Color{
		color.New(color.FgHiYellow),
		color.New(color.FgHiGreen),
		color.New(color.FgMagenta),
		color.New(color.FgHiBlue),
	}
	return colors[rand.Intn(len(colors))]
}

func colorizeAsciiArt(asciiArt string) string {
	coloredText := ""
	for _, char := range asciiArt {
		if char != ' ' && char != '\n' { // skip spaces and newlines
			randomColor := getRandomColor()
			coloredText += randomColor.Sprintf("%c", char)
		} else {
			coloredText += string(char)
		}
	}
	return coloredText
}

func greetUsers() {
	rand.Seed(uint64(time.Now().UnixNano()))
	fig := figure.NewFigure("MOVIES", "thick", true) // Other fonts alphabet, alligator2, thick, puffy, weird, wavy, twisted
	asciiArt := fig.String()

	// Apply color to each character in ASCII art
	coloredAsciiArt := colorizeAsciiArt(asciiArt)
	fmt.Println(coloredAsciiArt)
}
