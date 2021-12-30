package util

import (
	"fmt"
	"testing"
)

func TestId(t *testing.T) {
	id, err := NewIdWorker(0, 0, twepoch)
	if err != nil {
		//log.Error("NewIdWorker(0, 0) error(%v)", err)
		t.FailNow()
	}
	sid, err := id.NextId()
	if err != nil {
		//log.Error("id.NextId() error(%v)", err)
		t.FailNow()
	}

	fmt.Printf("snowflake id: %d \r\n", sid)
	sids, err := id.NextIds(10)
	if err != nil {
		//log.Error("id.NextId() error(%v)", err)
		t.FailNow()
	}
	fmt.Printf("snowflake ids: %v \n", sids)
	//log.Info("snowflake ids: %v", sids)
}
