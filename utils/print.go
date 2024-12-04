package utils

import (
	"fmt"

	cli "github.com/streamingfast/cli"
)

func PrintlnWarning(msg string) {
	fmt.Println(cli.WarningStyle.Render(msg))
}

func PrintWarning(msg string) {
	fmt.Print(cli.WarningStyle.Render(msg))
}

func PrintlnPurple(msg string) {
	fmt.Println(cli.PurpleStyle.Render(msg))
}

func PrintPurple(msg string) {
	fmt.Print(cli.PurpleStyle.Render(msg))
}

func PrintlnHeader(msg string) {
	fmt.Println(cli.HeaderStyle.Render(msg))
}

func PrintHeader(msg string) {
	fmt.Print(cli.HeaderStyle.Render(msg))
}

func PrintlnError(msg string) {
	fmt.Println(cli.ErrorStyle.Render(msg))
}

func PrintError(msg string) {
	fmt.Print(cli.ErrorStyle.Render(msg))
}
