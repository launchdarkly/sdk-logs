package main

import (
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/commands"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
)

const defaultCodesPath = "../../definition/codes/codes.json"
const defaultCodesSchemaPath = "../../definition/schema/ld_log_codes.json"
const defaultDocPath = "../../docs/codes.md"
const defaultFragmentPath = "../../docs/fragments"

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `Create and manage log specifiers
		USAGE:
			log <command>

		The default configuration of this command expects to be executed from its source directory.

		ENVIRONMENT VARIABLES:
			LOGDRKLY_CODES_PATH: Path to the codes JSON file to read/write code definitions.
				default: %s
			LOGDRKLY_SCHEMA_PATH: Path to the codes JSON schema file.
				default: %s
			LOGDRKLY_DOC_PATH: Generated documentation output file.
				default: %s
			LOGDRKLY_FRAGMENT_PATH: Path to find markdown fragments for generation.
				default: %s

		COMMANDS:
			new: Create a new system, class, or condition.
			deprecate: Deprecate a log code.
			supersede: Indicate a log code has been superseded.
			document: Generate markdown documentation for log codes.
			validate: Validate codes against the schema.
		`, defaultCodesPath, defaultCodesSchemaPath, defaultDocPath, defaultFragmentPath)
}

func getEnvWithDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func validateCodes(codesPath string, schemaPath string) {
	err := logs.ValidateCodes(codesPath, schemaPath)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {

	codePath := getEnvWithDefault("LOGDRKLY_CODES_PATH", defaultCodesPath)
	schemaPath := getEnvWithDefault("LOGDRKLY_SCHEMA_PATH", defaultCodesSchemaPath)
	docPath := getEnvWithDefault("LOGDRKLY_DOC_PATH", defaultDocPath)
	fragmentPath := getEnvWithDefault("LOGDRKLY_FRAGMENT_PATH", defaultFragmentPath)

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		validateCodes(codePath, schemaPath)
		commands.RunNewCommand(codePath, fragmentPath)
	case "deprecate":
		validateCodes(codePath, schemaPath)
		commands.RunDeprecateCommand(codePath)
	case "supersede":
		validateCodes(codePath, schemaPath)
		commands.RunSupersedeCommand(codePath)
	case "document":
		validateCodes(codePath, schemaPath)
		commands.RunDocumentCommand(codePath, fragmentPath, docPath)
	case "validate":
		validateCodes(codePath, schemaPath)
		fmt.Println("Codes matched schema.")
	default:
		fmt.Printf("Unrecognized command: %s\n", os.Args[1])
		usage()
	}

	os.Exit(0)
}
