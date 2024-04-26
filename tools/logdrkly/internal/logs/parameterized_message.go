package logs

import "errors"

type parserState int

var (
	stateCharacterEscape         = '\\'
	stateCharacterTokenStart     = '$'
	stateCharacterParameterStart = '{'
	stateCharacterParameterEnd   = '}'
)

var (
	parserStateIdle       parserState = 0
	parserStateEscape     parserState = 1
	parserMaybeTokenStart parserState = 2
	parserStateParameter  parserState = 3
)

type ParsedMessage struct {
	Message    string
	Parameters []string
}

func ParseMessage(messageString string) (ParsedMessage, error) {
	state := parserStateIdle

	parameterName := ""
	var parameters []string

	for _, char := range messageString {
		switch state {
		case parserStateIdle:
			switch char {
			case stateCharacterTokenStart:
				state = parserMaybeTokenStart
			case stateCharacterEscape:
				state = parserStateEscape
			default:
				continue
			}
		case parserStateEscape:
			state = parserStateIdle
		case parserMaybeTokenStart:
			switch char {
			case stateCharacterParameterStart:
				state = parserStateParameter
			default:
				continue
			}
		case parserStateParameter:
			switch char {
			case stateCharacterParameterEnd:
				state = parserStateIdle
				parameters = append(parameters, parameterName)
				parameterName = ""
			default:
				if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
					parameterName = parameterName + string(char)
					continue
				}
				return ParsedMessage{}, errors.New("invalid character in message parameter must be only ASCII letters")
			}
		}
	}
	switch state {
	case parserStateEscape:
		return ParsedMessage{}, errors.New("message string ended in escape character")
	case parserStateParameter:
		return ParsedMessage{}, errors.New("message string ended in an open parameter")
	}
	return ParsedMessage{
		Message:    messageString,
		Parameters: parameters,
	}, nil
}
