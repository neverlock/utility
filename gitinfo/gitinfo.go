package gitinfo

import (
	"io/ioutil"
	"os"
)

func Info(baseDir string, branch string) (string, error) {
	gitID := baseDir + ".git/refs/heads/" + branch
	if _, err := os.Stat(gitID); err != nil {
		return "This path don't have git commit id in " + gitID, err
	}

	id, err := ioutil.ReadFile(gitID)
	if err != nil {
		return "Can't open " + gitID, err
	}
	return string(id), nil

}
