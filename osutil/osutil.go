package osutil

import (
	"io"
	"os"
	"os/exec"
)

// Cheap hack, to get a file copy
func Copy(src, dest string) error {
	cmd := exec.Command("cp", "-a", src, dest)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	return cmd.Run()
}
