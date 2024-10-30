package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/charmbracelet/glamour"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/markdown"
)

func usage() {
	_, _ = fmt.Fprintln(os.Stderr, `Explains log codes

	USAGE:
		logdrkly explain <code>

	EXAMPLE:
		logdrkly explain 1:4:0
	`)
}

func RunExplainCommand(codesPath string) {
	if len(os.Args) < 3 {
		usage()
		return
	}
	_ = logs.WithCodes(codesPath, func(codes *logs.LdLogCodesJson) error {
		inputCode := os.Args[2]
		codeExpr := regexp.MustCompile(`[0-9]+:[0-9]+:[0-9]+`)
		matches := codeExpr.FindAllString(inputCode, -1)

		if len(matches) != 0 {
			for _, match := range matches {
				condition, present := codes.Conditions[match]
				if present {
					markdownCondition, err := markdown.GenMarkdownCondition(codes, "/tmp", logs.GetCode(condition))
					if err == nil {
						out, err := glamour.Render(markdownCondition, "dark")
						if err != nil {
							_, _ = fmt.Fprintf(os.Stderr, "Failed to render markdown: %s", err.Error())
							// Not treating this as an error.
						}
						fmt.Print(out)
					}
				}
			}
		} else {
			_, _ = fmt.Fprintln(os.Stderr, "Error code not in the correct format. Expected [0-9]+:[0-9]+:[0-9]+")
			usage()
		}

		return nil
	})
}
