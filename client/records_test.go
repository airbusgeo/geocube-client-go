package client

import (
	"reflect"
	"testing"
	"time"
)

func TestGroupBy(t *testing.T) {
	records := []*Record{
		{ID: "1", Time: time.Date(2022, 1, 1, 1, 1, 1, 1, time.Local)},
		{ID: "2", Time: time.Date(2022, 1, 1, 1, 1, 1, 1, time.Local)},
		{ID: "3", Time: time.Date(2022, 1, 2, 1, 1, 1, 1, time.Local)},
		{ID: "4", Time: time.Date(2022, 1, 2, 1, 1, 1, 1, time.Local)},
		{ID: "5", Time: time.Date(2022, 1, 1, 1, 1, 1, 1, time.Local)},
		{ID: "6", Time: time.Date(2022, 1, 3, 1, 1, 1, 1, time.Local)},
	}

	results := GroupBy(records, GroupByDayKey)
	expected := [][]string{{"1", "2", "5"}, {"3", "4"}, {"6"}}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("%v != %v", results, expected)
	}

}
