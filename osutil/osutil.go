package osutil

import (
	"io"
	"os"
	"os/exec"
	//"github.com/vbatts/golang2pkg/host"
)

/*
So copying files in bulk is not yet implemented and is left as an
exercise for the developer.

For now I'll hack. Later may implement a filepath.Walk, io.Copy madness
(remembering to mkdirs and set mtimes)
*/
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
