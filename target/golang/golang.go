package golang

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

//Run runs the target.
func (t *Target) Run() error {
	file, err := ioutil.TempFile(os.TempDir(), "*.go")
	if err != nil {
		return fmt.Errorf("could not create temp file: %w", err)
	}

	_, err = t.WriteTo(file)
	if err != nil {
		return fmt.Errorf("could not wrote temp file: %w", err)
	}

	var name = file.Name()
	file.Close()

	var cmd = exec.Command("go", "run", name)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
