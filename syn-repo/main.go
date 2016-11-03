package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func execCommand(dir string, name string, arg ...string) error {
	if len(dir) > 0 {
		os.Chdir(dir)
		fmt.Println(dir)
	}
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func pullGitRepo(dir string) {
	execCommand(dir, "git", "pull")
}

func pullGitRepoDirs(dir string) {
	if files, err := ioutil.ReadDir(dir); err == nil {
		for _, d := range files {
			if d.IsDir() {
				pullGitRepo(filepath.Join(dir, d.Name()))
			}
		}
	}
}

func main() {
	for {
		pullGitRepoDirs("/Users/jerry/Documents/ali/repo")
		pullGitRepoDirs("/Users/jerry/Documents/open-source")
		break
	}
}
