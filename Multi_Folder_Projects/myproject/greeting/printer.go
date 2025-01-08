package greeting

import "github.com/fatih/color"

// PrintColoredMessage prints a message in bold green
func PrintColoredMessage(message string) {
	color.Green(message)
}
