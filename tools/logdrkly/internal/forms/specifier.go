package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
)

func NewSpecifierForm(specifierType *logs.SpecifierType) *huh.Form {
	// Starting group decides which type of specifier to create.
	return huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().
			Title("Choose the type of specifier to create").
			Options(
				huh.NewOption("system", string(logs.SpecifierTypeSystem)),
				huh.NewOption("class", string(logs.SpecifierTypeClass)),
				huh.NewOption("condition", string(logs.SpecifierTypeCondition)),
			).Value((*string)(specifierType)),
	))
}
