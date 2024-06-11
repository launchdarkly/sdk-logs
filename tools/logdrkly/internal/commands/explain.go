package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/charmbracelet/glamour"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/markdown"
)

func RunExplainCommand(codesPath string) {
	_ = logs.WithCodes(codesPath, func(codes *logs.LdLogCodesJson) error {
		s := bufio.NewScanner(os.Stdin)
		codeExpr := regexp.MustCompile(`[0-9]+:[0-9]+:[0-9]+`)
		for s.Scan() {
			text := s.Text()
			matches := codeExpr.FindAllString(text, -1)

			log.Println(s.Text())
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
			}
		}

		return nil
	})
}
