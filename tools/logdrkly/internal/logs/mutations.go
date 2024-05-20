package logs

import (
	"errors"
	"fmt"

	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/collections"
)

func AddSystem(codes *LdLogCodesJson, name string, description string) error {
	err := ValidateSystemName(name, codes)
	if err != nil {
		return err
	}

	maxSpecifier := -1.0
	for _, system := range codes.Systems {
		if system.Specifier > maxSpecifier {
			maxSpecifier = system.Specifier
		}
	}
	newSpecifier := maxSpecifier + 1
	codes.Systems[name] = System{
		Description: description,
		Specifier:   newSpecifier,
	}
	return nil
}

func AddClass(codes *LdLogCodesJson, name string, description string) error {
	err := ValidateClassName(name, codes)
	if err != nil {
		return err
	}

	maxSpecifier := -1.0
	for _, system := range codes.Classes {
		if system.Specifier > maxSpecifier {
			maxSpecifier = system.Specifier
		}
	}
	newSpecifier := maxSpecifier + 1
	codes.Classes[name] = Class{
		Description: description,
		Specifier:   newSpecifier,
	}
	return nil
}

func AddCondition(codes *LdLogCodesJson, className string, systemName string, condName string, description string, message Message) (Condition, error) {
	system, systemPresent := codes.Systems[systemName]

	if !systemPresent {
		return Condition{}, fmt.Errorf("the system class does not exist. Please choose an existing system or create a new system")
	}

	class, classPresent := codes.Classes[className]

	if !classPresent {
		return Condition{}, fmt.Errorf("the specified class does not exist. Please choose an existing class or create a new class")
	}

	conditions := collections.MapFilter(codes.Conditions, func(condition Condition) bool {
		return condition.System == system.Specifier && condition.Class == class.Specifier
	})

	maxSpecifier := -1.0
	for _, condition := range conditions {
		if condition.Specifier > maxSpecifier {
			maxSpecifier = condition.Specifier
		}
	}
	newSpecifier := maxSpecifier + 1

	condition := Condition{
		Name:        condName,
		Description: description,
		Specifier:   newSpecifier,
		Class:       class.Specifier,
		System:      system.Specifier,
		Message:     message,
	}

	code := GetCode(condition)

	_, present := codes.Conditions[code]
	if present {
		return Condition{}, fmt.Errorf("condition code already exists. Please choose a new code or the existing condition")
	}

	codes.Conditions[code] = condition

	return condition, nil
}

func DeprecateCode(codes *LdLogCodesJson, code string, reason string) error {
	condition, present := codes.Conditions[code]
	if !present {
		return errors.New("cannot deprecate a condition which does not exist")
	}
	deprecated := true
	condition.Deprecated = &deprecated
	condition.DeprecatedReason = &reason
	codes.Conditions[code] = condition
	return nil
}

func SupersedeCode(codes *LdLogCodesJson, code string, replacementCode string, reason string) error {
	condition, present := codes.Conditions[code]
	if !present {
		return errors.New("cannot deprecate a condition which does not exist")
	}
	condition.Superseded = &replacementCode
	condition.SupersededReason = &reason
	codes.Conditions[code] = condition
	return nil
}
