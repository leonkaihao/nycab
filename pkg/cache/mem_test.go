package cache

import (
	"testing"
	"time"
)

func Test_mem_suite(t *testing.T) {
	cache, _ := NewMemCache()
	tm11, _ := time.Parse("2006-01-02", "2011-11-11")
	// create a value
	cache.WithData("id1", tm11, 10)
	count, _ := cache.GetDataSingle("id1", tm11)
	if count != 10 {
		t.Error("failed to get data")
	}
	// update a value
	cache.WithData("id1", tm11, 11)
	count, _ = cache.GetDataSingle("id1", tm11)
	if count != 11 {
		t.Error("failed to update data")
	}
	// create another
	cache.WithData("id2", tm11, 12)
	info, _ := cache.GetDataMulti([]string{"id1", "id2"}, tm11)
	if len(info) != 2 {
		t.Error("failed to compare data size")
	}
	if info[0].Count != 11 || info[1].Count != 12 {
		t.Error("data not correct")
	}
	// clear data
	cache.Clear()
	info, _ = cache.GetDataMulti([]string{"id1", "id2"}, tm11)
	if len(info) != 0 {
		t.Error("failed to clean cache")
	}
}
