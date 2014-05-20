// +build linux

package linux

import (
	"github.com/shirou/gopsutil/structs"
	"github.com/shirou/gopsutil/utils"

	"strings"
)

func NetIOCounters(pernic bool) ([]structs.NetIOCountersStat, error) {
	filename := "/proc/net/dev"
	lines, err := utils.ReadLines(filename)
	if err != nil {
		return nil, err
	}

	statlen := len(lines) - 1

	ret := make([]structs.NetIOCountersStat, 0, statlen)

	for _, line := range lines[2:] {
		fields := strings.Fields(line)
		if fields[0] == "" {
			continue
		}
		nic := structs.NetIOCountersStat{
			Name:        strings.Trim(fields[0], ":"),
			BytesRecv:   utils.MustParseUint64(fields[1]),
			Errin:       utils.MustParseUint64(fields[2]),
			Dropin:      utils.MustParseUint64(fields[3]),
			BytesSent:   utils.MustParseUint64(fields[9]),
			PacketsSent: utils.MustParseUint64(fields[10]),
			Errout:      utils.MustParseUint64(fields[11]),
			Dropout:     utils.MustParseUint64(fields[12]),
		}
		ret = append(ret, nic)
	}
	return ret, nil
}
