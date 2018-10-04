package cgroups

import "mydocker/cgroups/subsystems"

type CgroupManager struct {
}

func NewCgroupManager(path string) *CgroupManager {

}

func (c *CgroupManager) Apply(pid int) error {

}

func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {

}

func (c *CgroupManager) Destory() error {

}
