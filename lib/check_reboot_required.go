package reboot_required

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

const rebootRequiredFile = "/var/run/reboot-required"

func Do() {
	ckr := run(os.Args[1:])
	ckr.Name = "RebootRequired"
	ckr.Exit()
}

var opts struct {
	File string `short:"f" long:"file" default:"/var/run/reboot-required" description:"reboot required file name"`
}

func run(args []string) *checkers.Checker {
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		os.Exit(1)
	}

	_, err = os.Stat(opts.File)

	msg := ""
	checkSt := checkers.OK
	switch {
	case os.IsNotExist(err):
		msg = "reboot is not required"
		checkSt = checkers.OK
	case err == nil:
		msg = "reboot is required"
		checkSt = checkers.WARNING
	default:
		msg = fmt.Sprintf("Failed to check: %s", err)
		checkSt = checkers.UNKNOWN
	}

	return checkers.NewChecker(checkSt, msg)
}
