package worm

import (
	"testing"
	. "wego/worm/conf"
)

func TestWeightSlave(t *testing.T)  {
	db, err := OpenDb()
	if err != nil {
		t.Error(err)
		return
	}

	eng, err := NewEngine("mysql", db)
	if err != nil {
		t.Error(err)
		return
	}

	eng.AddSlave(db, "db2", 1)
	eng.AddSlave(db, "db3", 2)
	ret := getSlaveByWeight(eng.slaves)
	t.Log(ret.db_name)
}
