package exec

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func ExecCmd(cmd string, args ...string) error {
	_, _, err := ExecCmdOut(cmd, args...)
	return err
}

func ExecCmdOut(cmd string, args ...string) (string, string, error) {
	return ExecCmdOutWithStdin("", cmd, true, args...)
}

func ExecCmdOutWithStdin(stdinContent string, cmd string, isLoggingEnabled bool, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var stdin *bytes.Buffer = bytes.NewBuffer([]byte{})
	if stdinContent != "" {
		stdin = bytes.NewBuffer([]byte(stdinContent))
	}

	cmdErr := execCmdStd(context.Background(), nil, stdin, cmd, &stdout, &stderr, isLoggingEnabled, args...)

	return stdout.String(), stderr.String(), cmdErr
}

func execCmdStd(ctx context.Context, env map[string]string, stdin io.Reader, cmd string, stdout io.Writer, stderr io.Writer, isLoggingEnabled bool, args ...string) error {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return fmt.Errorf("could not find executable %s on path: %w", cmd, err)
	}
	c := exec.CommandContext(ctx, cmd, clearItemsFromList(args)...)

	if isLoggingEnabled {
		c.Stdout = NewTeeWriter(os.Stdout, stdout)
		c.Stderr = NewTeeWriter(os.Stderr, stderr)
	} else {
		c.Stdout = stdout
		c.Stderr = stderr
	}
	c.Stdin = stdin

	c.Env = make([]string, 0)
	for key, value := range env {
		c.Env = append(c.Env, key+"="+value)
	}

	// add os environment
	for _, envPair := range os.Environ() {
		parts := strings.Split(envPair, "=")
		key := parts[0]
		if _, contains := env[key]; !contains {
			c.Env = append(c.Env, envPair)
		}
	}

	cmdString := cmd + " " + strings.Join(args, " ")
	fmt.Println("Execute: " + cmdString)

	return c.Run()
}

func clearItemsFromList(args []string) []string {
	var r []string
	for _, str := range args {
		if str != "" {
			r = append(r, strings.TrimSpace(str))
		}
	}
	return r
}
