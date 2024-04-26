package commands

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/forms"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/markdown"
)

func RunNewCommand(codePath string, fragmentPath string) {
	var specifierType logs.SpecifierType
	form := forms.NewSpecifierForm(&specifierType)
	err := form.Run()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error specifier form", err.Error())
		return
	}
	switch specifierType {
	case logs.SpecifierTypeSystem:
		runNewSystemCommand(codePath)
	case logs.SpecifierTypeClass:
		runNewClassCommand(codePath)
	case logs.SpecifierTypeCondition:
		runNewConditionCommand(codePath, fragmentPath)
	}
}

func runNewClassCommand(codePath string) {
	err := logs.UpdateCodes(codePath, func(codes *logs.LdLogCodesJson) error {
		var params forms.ClassFormData
		form := forms.NewClassForm(codes, &params)
		err := form.Run()
		if err != nil {
			return fmt.Errorf("error running new system form: %w", err)
		}
		err = logs.AddClass(codes, params.Name, params.Description)
		return err
	})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func runNewSystemCommand(codePath string) {
	err := logs.UpdateCodes(codePath, func(codes *logs.LdLogCodesJson) error {
		var params forms.SystemFormData
		form := forms.NewSystemForm(codes, &params)
		err := form.Run()
		if err != nil {
			return fmt.Errorf("error running new system form: %w", err)
		}
		err = logs.AddSystem(codes, params.Name, params.Description)
		return err
	})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func runNewConditionCommand(codePath string, fragmentPath string) {
	err := logs.UpdateCodes(codePath, func(codes *logs.LdLogCodesJson) error {
		var params forms.ConditionFormData
		form := forms.NewConditionFormPart1(codes, &params)
		err := form.Run()

		if err != nil {
			return fmt.Errorf("error running new condition form part 1: %w", err)
		}

		form = forms.NewConditionFormPart2(codes, &params)
		err = form.Run()
		if err != nil {
			return fmt.Errorf("error running new condition form part 2: %w", err)
		}

		parameters, err := logs.ParseMessage(params.MessageString)
		if err != nil {
			return fmt.Errorf("bad message string: %w", err)
		}

		message := logs.Message{
			Parameters:    map[string]string{},
			Parameterized: params.MessageString,
		}

		if len(parameters.Parameters) != 0 {
			var paramParams forms.ParameterFormData
			parametersForm := forms.NewMessageForm(codes, parameters.Parameters, &paramParams)
			err = parametersForm.Run()
			if err != nil {
				return fmt.Errorf("error running parameters form: %w", err)
			}
			for key, value := range paramParams.Descriptions {
				message.Parameters[key] = *value
			}
		}

		condition, err := logs.AddCondition(codes, params.Class, params.System, params.Name, params.Description, message)
		if err != nil {
			return err
		}

		fmt.Printf("The \"%s\" condition has been added.\n", params.Name)
		markdownCondition, err := markdown.GenMarkdownCondition(codes, fragmentPath, logs.GetCode(condition))
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Could not generate markdown preview.")
		}

		out, err := glamour.Render(markdownCondition, "dark")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to render markdown: %s", err.Error())
			// Not treating this as an error.
		}
		fmt.Print(out)

		return nil
	})
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}
