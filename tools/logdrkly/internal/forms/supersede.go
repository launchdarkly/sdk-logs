package forms

import (
	"errors"

	"github.com/charmbracelet/huh"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
)

type SupersedeFormData struct {
	SupersededCode string
	NewCode        string
	Reason         string
}

func NewSupersedeForm(codes *logs.LdLogCodesJson, deprecate *SupersedeFormData) *huh.Form {
	// Starting group decides which type of specifier to create.
	return huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Please enter the code to supersede?").
			Value(&deprecate.SupersededCode).
			Validate(func(s string) error {
				if !logs.ValidateCode(s) {
					return errors.New("the code was not in the correct format")
				}
				_, present := codes.Conditions[s]

				if !present {
					return errors.New("could not find an existing entry matching the code")
				}
				return nil
			}),
		huh.NewInput().
			Title("Please enter the code that is superseding it?").
			Value(&deprecate.NewCode).
			Validate(func(s string) error {
				if !logs.ValidateCode(s) {
					return errors.New("the code was not in the correct format")
				}
				_, present := codes.Conditions[s]

				if !present {
					return errors.New("could not find an existing entry matching the code")
				}
				return nil
			}),
		huh.NewInput().
			Title("What is the reason the code is being superseded?").
			Value(&deprecate.Reason),
	))
}
