package subsystems

type CpusetSubsystem struct {
}

func (s *CpusetSubsystem) Set(cgroupPath string, res *ResourceConfig) error {

}

func (s *CpusetSubsystem) Remove(cgroupPath string) error {

}

func (s *CpusetSubsystem) Apply(cgroupPath string, pid int) error {

}

func (s *CpusetSubsystem) Name() string {

}
