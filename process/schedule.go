package process

import (
	"os/exec"

	"github.com/vshaveyko/test-go-daemon/dlog"
)

func scheduleProcessee(pipelines []Pipeline) {

	for _, p := range pipelines {
		cmd := p.toShellCmd()

		err := execCmd(cmd)

		if err != nil {
			dlog.Stdlog.Printf("Encountered err: %s", err)
		}
	}

}

func execCmd(cmdDef []string) error {

	cmd := exec.Command(cmdDef[0], cmdDef[1:]...)

	dlog.Stdlog.Printf("Running command and waiting for it to finish...")

	out, err := cmd.CombinedOutput()

	if err != nil {
		dlog.Stdlog.Printf("Encountered err: %s, %s, %s", cmdDef, out, err)
	} else {
		dlog.Stdlog.Printf("Executed command: %s, %s", cmdDef, out)
	}

	return err

}
