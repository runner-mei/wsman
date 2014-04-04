package winrm

import (
	"bytes"
	"github.com/dylanmei/packer-communicator-winrm/winrm"
	"log"
	"strings"
)

func Cmd(shell *winrm.Shell, command string) (string, error) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	shell.Stdout = stdout
	shell.Stderr = stderr

	log.Println("starting cmd:", command)

	c, err := shell.NewCommand(command)

	if err != nil {
		return "", err
	}

	if _, err = c.Receive(); err != nil {
		return "", err
	}

	if stderr.Len() > 0 {
		log.Println("cmd stderr: %s", stderr.String())
	}

	return strings.Trim(stdout.String(), " \r\n"), nil
}
