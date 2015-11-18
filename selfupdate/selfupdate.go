package selfupdate

import (
	"fmt"
	"github.com/sanbornm/go-selfupdate/selfupdate"
)

func SelfUpdate(version string, url string, appName string) error {
	updater := &selfupdate.Updater{
		CurrentVersion: version,
		ApiURL:         url,
		BinURL:         url,
		DiffURL:        url,
		Dir:            "update-tmp",
		CmdName:        appName, // app name
	}
	if updater != nil {
		fmt.Printf("[%s][%s] Checking for new versions...", appName, version)
		err := updater.BackgroundRun()
		if err != nil {
			return err
		}
	}
	return nil
}
