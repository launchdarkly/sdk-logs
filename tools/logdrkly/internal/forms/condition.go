package forms

import (
	"github.com/charmbracelet/huh"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/collections"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
)

type ConditionFormData struct {
	Description   string
	Name          string
	MessageString string
	System        string
	Class         string
}

func NewConditionFormPart1(codes *logs.LdLogCodesJson, condition *ConditionFormData) *huh.Form {
	var systemOptions []huh.Option[string]
	var classOptions []huh.Option[string]

	// It is nice if these are always presented in
	collections.MapForEachOrdered(codes.Systems, func(systemName string, _ logs.System) {
		systemOptions = append(systemOptions, huh.NewOption(systemName, systemName))
	})

	collections.MapForEachOrdered(codes.Classes, func(className string, _ logs.Class) {
		classOptions = append(classOptions, huh.NewOption(className, className))
	})

	return huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().Title("Select system:").Options(systemOptions...).Value(&condition.System),
		huh.NewSelect[string]().Title("Select class:").Options(classOptions...).Value(&condition.Class),
	))
}

func NewConditionFormPart2(codes *logs.LdLogCodesJson, condition *ConditionFormData) *huh.Form {
	return huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("The name of the condition?").
			Value(&condition.Name).
			Validate(func(s string) error {
				return logs.ValidateConditionName(s, codes.Systems[condition.System].Specifier,
					codes.Classes[condition.Class].Specifier, codes)
			}),
		huh.NewInput().
			Title("Please describe the condition?").
			Value(&condition.Description),
		huh.NewInput().
			Title("Parameterized message string for the condition:").
			Value(&condition.MessageString).
			Validate(func(s string) error {
				return logs.ValidateParameterizedMessageString(condition.MessageString)
			}),
	))
}
