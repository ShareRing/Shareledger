package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const (
	WorkDir = "./build/"
	KeyRing = "os"
	NSD     = WorkDir + "shareledger"
	NSCLI   = WorkDir + "slcli"
	chainID = "SHR-VoyagerNet"
)

var (
	makePath          string    = "/usr/bin/make"
	dockerComposePath           = "/usr/local/bin/docker-compose"
	output            io.Writer = os.Stdout
)

func buildBinaries() error {
	_, err := execCmdAndWait(makePath, "clean", "build", "build-docker")
	if err != nil {
		fmt.Println("Build failed. error: ", err)
	}
	return err
}

func dockerComposeUp() error {
	return execCmd(dockerComposePath, []string{"up"})
}

func dockerComposeDown() error {
	return execCmd(dockerComposePath, []string{"down"})
}

func createTestnet() error {
	_, err := execCmdAndWait(NSD, "testnet", chainID, "-o", WorkDir, "--keyring-backend", KeyRing)
	if err != nil {
		return err
	}
	return nil
}

func execCmd(name string, arguments []string) error {
	cmd := exec.Command(name, arguments...)
	err := writeoutput(cmd, nil)
	if err != nil {
		return err
	}
	return cmd.Start()
}

func execCmdAndWait(name string, arguments ...string) (string, error) {
	cmd := exec.Command(name, arguments...)

	var output strings.Builder

	captureOutput := func(s string) {
		output.WriteString(s)
		output.WriteRune('\n')
	}
	err := writeoutput(cmd, captureOutput)

	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func writeoutput(cmd *exec.Cmd, captureOutput func(string)) error {
	stderrReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	go outputHandler("stderr", stderrReader, captureOutput)
	go outputHandler("stdout", stdoutReader, captureOutput)
	return nil
}

func outputHandler(prefix string, reader io.Reader, captureOutput func(string)) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Fprintf(output, "%s | %s\n", prefix, s)
		if captureOutput != nil {
			captureOutput(s)
		}
	}
}
