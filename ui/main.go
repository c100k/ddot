package ui

import "fmt"

func PrintErr(err error) {
	fmt.Printf("❌ %v\n", err)
}

func Print(emoji string, msg string) {
	fmt.Printf("%s %s\n", emoji, msg)
}

func PrintSuccess(msg string) {
	fmt.Printf("✅ %s\n", msg)
}

func Prompt(emoji string, msg string) {
	fmt.Printf("%s %s => ", emoji, msg)
}
