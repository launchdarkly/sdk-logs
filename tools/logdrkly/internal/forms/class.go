package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
)

type ClassFormData struct {
	Description string
	Name        string
}

func NewClassForm(codes *logs.LdLogCodesJson, cls *ClassFormData) *huh.Form {
	// Starting group decides which type of specifier to create.
	return huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("The name of the class?").
			Value(&cls.Name).
			Validate(func(s string) error {
				return logs.ValidateClassName(s, codes)
			}),
		huh.NewInput().
			Title("Please describe the class?").
			Value(&cls.Description),
	))
}
