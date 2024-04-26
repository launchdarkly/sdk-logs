package logs

import (
	"fmt"
	"regexp"
)

const specifierName = "^[a-zA-Z]+$"

func ValidSpecifierName(s string) bool {
	matched, err := regexp.MatchString(specifierName, s)
	if err != nil {
		return false
	}
	return matched
}

type SpecifierType string

var (
	SpecifierTypeClass     SpecifierType = "class"
	SpecifierTypeSystem    SpecifierType = "system"
	SpecifierTypeCondition SpecifierType = "condition"
)

func GetCode(condition Condition) string {
	return fmt.Sprintf("%d:%d:%d", int(condition.System), int(condition.Class), int(condition.Specifier))
}
