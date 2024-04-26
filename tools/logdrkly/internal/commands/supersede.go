package commands

import (
	"fmt"
	"os"

	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/forms"
	"github.com/launchdarkly/sdk-logs/tools/logdrkly/internal/logs"
)

func RunSupersedeCommand(codePath string) {
	var supersedeParams forms.SupersedeFormData
	err := logs.UpdateCodes(codePath, func(codes *logs.LdLogCodesJson) error {
		form := forms.NewSupersedeForm(codes, &supersedeParams)
		err := form.Run()
		if err != nil {
			return err
		}

		err = logs.SupersedeCode(codes, supersedeParams.SupersededCode, supersedeParams.NewCode, supersedeParams.Reason)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		return
	}
	fmt.Println("Superseded code")
}
