package cmdrunner

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/specterops/bloodhound/log"
	"github.com/specterops/bloodhound/packages/go/stbernard/environment"
)

var (
	ErrNonZeroExit = errors.New("non-zero exit status")
)

// CmdModifier describes a function that mutates a generated *exec.Cmd before running it
type CmdModifier func(*exec.Cmd)

// Run runs a command with ars and environment variables set at a specified path
//
// The CmdModifiers parameter is an optional list of modifying functions that can alter the generated *exec.Cmd after default setup.
//
// If debug log level is set globally, command output will be combined and sent to os.Stderr.
func Run(command string, args []string, path string, env environment.Environment, cmdModifiers ...CmdModifier) error {
	var (
		exitErr error

		cmdstr       = command + " " + args[0]
		cmd          = exec.Command(command, args...)
		debugEnabled = log.GlobalAccepts(log.LevelDebug)
	)

	cmd.Env = env.Slice()
	cmd.Dir = path

	if debugEnabled {
		cmd.Stdout = os.Stderr
		cmd.Stderr = os.Stderr
		cmdstr = command + " " + strings.Join(args, " ")
	}

	// If we got any cmdModifiers, apply them in order
	if len(cmdModifiers) > 0 {
		for _, modifier := range cmdModifiers {
			modifier(cmd)
		}
	}

	log.Infof("Running %s for %s", cmdstr, path)

	err := cmd.Run()
	if _, ok := err.(*exec.ExitError); ok {
		exitErr = ErrNonZeroExit
	} else if err != nil {
		return fmt.Errorf("%s: %w", cmdstr, err)
	}

	log.Infof("Finished %s for %s", cmdstr, path)

	return exitErr
}