package cli

import (
	"fmt"
	"github.com/charmbracelet/huh"
)

func RunConfirmForm(title string) (bool, error) {
	confirmOverwrite := true

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Value(&confirmOverwrite).
				Affirmative("Yes").
				Negative("No"),
		),
	)

	if err := form.Run(); err != nil {
		return false, fmt.Errorf("error running form: %w", err)
	}

	return confirmOverwrite, nil
}
