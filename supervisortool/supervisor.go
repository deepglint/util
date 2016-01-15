package supervisortool

import (
	"errors"

	"github.com/deepglint/glog"
	"github.com/deepglint/go-supervisor/supervisor"
	"github.com/deepglint/util/process"
	"github.com/deepglint/util/systool"
)

func ListSupervisorProcess(client supervisor.Client) (pids []int, err error) {
	var finalError bool
	superProcesses, err := client.GetAllProcessInfo()
	if err != nil {
		return
	}

	for _, superProcess := range superProcesses {
		glog.Infoln(superProcess)
		if superProcess.PID == 0 {
			glog.Warningf("fatal process %s", superProcess.Name)
			continue
		}
		childProcesses, err := process.FindPidsFromPPid(int(superProcess.PID))
		if err != nil {
			glog.Errorln(err)
			finalError = true
			continue
		}

		for _, childProcess := range childProcesses {
			glog.Infoln(childProcess)
			pids = append(pids, childProcess.Pid())
		}
	}

	if finalError {
		err = errors.New("failed to list one of processes")
	}

	return
}

func StopSupervisor() (err error) {
	//
	_, err = systool.CmdOut("/usr/bin/sudo", "/usr/bin/killall", "supervisord")

	return
}

func StartSupervisor() (err error) {
	_, err = systool.CmdOut("/usr/bin/sudo", "/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf")

	return
}
