package util

import (
	"fmt"
	"testing"
)

func TestId(t *testing.T) {
	id, err := NewIdWorker(0, twepoch)
	if err != nil {
		t.FailNow()
	}
	sid, err := id.NextId()
	if err != nil {
		t.FailNow()
	}

	fmt.Printf("snowflake id: %d \r\n", sid)
	sids, err := id.NextIds(10)
	if err != nil {
		t.FailNow()
	}
	fmt.Printf("snowflake ids: %v \n", sids)
}
