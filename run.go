package main

import (
	log "github.com/Sirupsen/logrus"
	"mydocker/cgroups"
	"mydocker/cgroups/subsystems"
	"mydocker/container"
	"os"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}

func Run(tty bool, comArray []string, res *subsystems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}

	if err := parent.Start(); err != nil {
		log.Errorf(err)
	}

	cgroupMmanager := cgroups.CgroupManager("mydocker-cgroup")
	defer cgroupMmanager.Destory()

	cgroupMmanager.Set(res)

	cgroupMmanager.Apply(parent.Process.Pid)

	sendInitCommand(comArray, writePipe)

	parent.Wait()
}
