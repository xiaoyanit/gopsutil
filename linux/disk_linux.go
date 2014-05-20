// +build linux

package linux

import (
	"github.com/shirou/gopsutil/structs"
	"github.com/shirou/gopsutil/utils"

	"strings"
	"unicode"
)

const (
	SECTOR_SIZE = 512
)

// Get disk partitions.
// should use setmntent(3) but this implement use /etc/mtab file
func DiskPartitions(all bool) ([]structs.DiskPartitionStat, error) {

	filename := "/etc/mtab"
	lines, err := utils.ReadLines(filename)
	if err != nil {
		return nil, err
	}

	ret := make([]structs.DiskPartitionStat, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)
		d := structs.DiskPartitionStat{
			Mountpoint: fields[1],
			Fstype:     fields[2],
			Opts:       fields[3],
		}
		ret = append(ret, d)
	}

	return ret, nil
}

func DiskIOCounters() (map[string]structs.DiskIOCountersStat, error) {
	// determine partitions we want to look for
	filename := "/proc/partitions"
	lines, err := utils.ReadLines(filename)
	if err != nil {
		return nil, err
	}
	partitions := make([]string, 0, len(lines)-2)

	for _, line := range lines[2:] {
		fields := strings.Fields(line)
		name := []rune(fields[3])

		if unicode.IsDigit(name[len(name)-1]) {
			partitions = append(partitions, fields[3])
		} else {
			// http://code.google.com/p/psutil/issues/detail?id=338
			lenpart := len(partitions)
			if lenpart == 0 || strings.HasPrefix(partitions[lenpart-1], fields[3]) {
				partitions = append(partitions, fields[3])
			}
		}
	}

	filename = "/proc/diskstats"
	lines, err = utils.ReadLines(filename)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]structs.DiskIOCountersStat, 0)
	for _, line := range lines {
		fields := strings.Fields(line)
		name := fields[2]
		reads := utils.MustParseUint64(fields[3])
		rbytes := utils.MustParseUint64(fields[5])
		rtime := utils.MustParseUint64(fields[6])
		writes := utils.MustParseUint64(fields[7])
		wbytes := utils.MustParseUint64(fields[9])
		wtime := utils.MustParseUint64(fields[10])
		if utils.StringContains(partitions, name) {
			d := structs.DiskIOCountersStat{
				Name:       name,
				ReadBytes:  rbytes * SECTOR_SIZE,
				WriteBytes: wbytes * SECTOR_SIZE,
				ReadCount:  reads,
				WriteCount: writes,
				ReadTime:   rtime,
				WriteTime:  wtime,
			}
			ret[name] = d

		}
	}
	return ret, nil
}
