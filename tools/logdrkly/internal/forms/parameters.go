package forms

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
)

type ParameterFormData struct {
	Descriptions map[string]*string
}

func NewMessageForm(codes *logs.LdLogCodesJson, parameterNames []string, message *ParameterFormData) *huh.Form {
	message.Descriptions = make(map[string]*string)

	var inputs []huh.Field

	for _, parameterName := range parameterNames {
		var parameterOutput string
		message.Descriptions[parameterName] = &parameterOutput
		inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter description for the \"%s\" parameter:", parameterName)).Value(&parameterOutput))
	}

	// Starting group decides which type of specifier to create.
	return huh.NewForm(huh.NewGroup(inputs...))
}
