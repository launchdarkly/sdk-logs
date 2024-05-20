package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
)

type SystemFormData struct {
	Description string
	Name        string
}

func NewSystemForm(codes *logs.LdLogCodesJson, system *SystemFormData) *huh.Form {
	// Starting group decides which type of specifier to create.
	return huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("The name of the system?").
			Value(&system.Name).
			Validate(func(s string) error {
				return logs.ValidateSystemName(s, codes)
			}),
		huh.NewInput().
			Title("Please describe the system?").
			Value(&system.Description),
	))
}
