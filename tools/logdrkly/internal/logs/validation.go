package logs

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/collections"
	"github.com/xeipuuv/gojsonschema"
)

const codeFormat = "^[0-9]+:[0-9]+:[0-9]+$"

func validateName(name string, specifierType SpecifierType, present bool) error {
	if present {
		return fmt.Errorf("%s name already exists. Please choose a new name or use the existing specifier", specifierType)
	}
	if !ValidSpecifierName(name) {
		return fmt.Errorf("the %s name must be composed of only upper and lowercase ASCII letters and may not be empty [a-zA-Z]+", specifierType)
	}
	return nil
}

func ValidateSystemName(name string, codes *LdLogCodesJson) error {
	_, present := codes.Systems[name]
	return validateName(name, SpecifierTypeSystem, present)
}

func ValidateClassName(name string, codes *LdLogCodesJson) error {
	_, present := codes.Classes[name]
	return validateName(name, SpecifierTypeSystem, present)
}

func ValidateConditionName(name string, system float64, class float64, codes *LdLogCodesJson) error {
	conditions := collections.MapFilter(codes.Conditions, func(condition Condition) bool {
		return condition.System == system && condition.Class == class
	})
	_, _, present := collections.MapFind(conditions, func(s string, condition Condition) bool {
		return condition.Name == name
	})
	return validateName(name, SpecifierTypeCondition, present)
}

func ValidateParameterizedMessageString(message string) error {
	_, err := ParseMessage(message)
	return err
}

func ValidateCode(code string) bool {
	matched, err := regexp.MatchString(codeFormat, code)
	if err != nil {
		return false
	}
	return matched
}

func canonicalPath(path string) string {
	dir, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("could not access working directory: %w", err))
		os.Exit(1)
	}
	return fmt.Sprintf("file://%s/%s", dir, path)
}

func ValidateCodes(codesPath string, schemaPath string) error {
	schemaLoader := gojsonschema.NewReferenceLoader(canonicalPath(codesPath))
	documentLoader := gojsonschema.NewReferenceLoader(canonicalPath(schemaPath))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}
	if result.Valid() {
		return nil
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "The document is not valid. see errors :")
		for _, desc := range result.Errors() {
			_, _ = fmt.Fprintf(os.Stderr, "- %s\n", desc)
		}
		return errors.New("error processing codes json file")
	}
}
