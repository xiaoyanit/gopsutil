// +build linux

package test

import (
	"fmt"
	"testing"

	. "github.com/shirou/gopsutil"
	"github.com/shirou/gopsutil/structs"
)

func TestLoad(t *testing.T) {
	v, err := LoadAvg()
	if err != nil {
		t.Errorf("error %v", err)
	}

	empty := &structs.LoadAvgStat{}
	if v == empty {
		t.Errorf("error load: %v", v)
	}
}

func TestLoadAvgStat_String(t *testing.T) {
	v := structs.LoadAvgStat{
		Load1:  10.1,
		Load5:  20.1,
		Load15: 30.1,
	}
	e := `{"load1":10.1,"load5":20.1,"load15":30.1}`
	if e != fmt.Sprintf("%v", v) {
		t.Errorf("LoadAvgStat string is invalid: %v", v)
	}
}
