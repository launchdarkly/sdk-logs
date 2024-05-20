package commands

import (
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/markdown"
)

func RunDocumentCommand(codePath string, fragmentPath string, outPath string) {
	err := logs.WithCodes(codePath, func(codes *logs.LdLogCodesJson) error {
		return markdown.GenerateMarkdown(codes, fragmentPath, outPath)
	})

	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
	} else {
		fmt.Printf("Generated documentation in %s", outPath)
	}
}
