package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubsystem struct {
}

func (s *MemorySubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysGrpPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err != nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(path.Join(subsysGrpPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
				return fmt.Errorf("set cgroup memroy fail %v", err)
			}
		}
		return nil
	} else {
		return err
	}
}

func (s *MemorySubsystem) Remove(cgroupPath string) error {
	if subsysCgrpPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		return os.Remove(subsysCgrpPath)
	} else {
		return err
	}
}

func (s *MemorySubsystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %s error: %v", cgroupPath, err)
	}
}

func (s *MemorySubsystem) Name() string {
	return "memory"
}
