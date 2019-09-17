package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/waytkheming/godojo/dojo7/kadai2/waytkheming/cli"
)

func TestCLI_Run(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	c := cli.NewCLI(outStream, errStream)
	args := strings.Split("convert ./../testdata", " ")
	exitCode := c.Run(args)

	if exitCode != cli.ExitCodeOK {
		t.Errorf("failed cli run, exit_code: %d", exitCode)
	}

	if errStream.Len() > 0 {
		t.Errorf("failed cli run, output: %q", errStream.String())
	}
}
