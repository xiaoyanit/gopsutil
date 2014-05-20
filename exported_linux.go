// +build linux

package gopsutil

import (
	platform "github.com/shirou/gopsutil/linux"
	"github.com/shirou/gopsutil/structs"
	"runtime"
)

func CPUTimes(percpu bool) ([]structs.CPUTimesStat, error) {
	return platform.CPUTimes(percpu)
}

func CPUInfo() ([]structs.CPUInfoStat, error) {
	return platform.CPUInfo()
}
func DiskPartitions(all bool) ([]structs.DiskPartitionStat, error) {
	return platform.DiskPartitions(all)
}
func DiskIOCounters() (map[string]structs.DiskIOCountersStat, error) {
	return platform.DiskIOCounters()
}
func HostInfo() (*structs.HostInfoStat, error) {
	return platform.HostInfo()
}
func BootTime() (uint64, error) {
	return platform.BootTime()
}
func Users() ([]structs.UserStat, error) {
	return platform.Users()
}
func LoadAvg() (*structs.LoadAvgStat, error) {
	return platform.LoadAvg()
}
func VirtualMemory() (*structs.VirtualMemoryStat, error) {
	return platform.VirtualMemory()
}
func SwapMemory() (*structs.SwapMemoryStat, error) {
	return platform.SwapMemory()
}
func NetIOCounters(pernic bool) ([]structs.NetIOCountersStat, error) {
	return platform.NetIOCounters(pernic)
}
func NewProcess(pid int) (*structs.Process, error) {
	return platform.NewProcess(pid)
}
func Pids() ([]int32, error) {
	return platform.Pids()
}
func PidExists(pid int32) (bool, error) {
	return platform.PidExists(pid)
}


func CPUCounts(logical bool) (int, error) {
	return runtime.NumCPU(), nil
}
