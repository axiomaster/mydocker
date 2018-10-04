package subsystems

type CpuSubsystem struct {
}

func (s *CpuSubsystem) Set(cgroupPath string, res *ResourceConfig) error {

}

func (s *CpuSubsystem) Remove(cgroupPath string) error {

}

func (s *CpuSubsystem) Apply(cgroupPath string, pid int) error {

}

func (s *CpuSubsystem) Name() string {

}
