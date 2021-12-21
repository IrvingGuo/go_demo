package util_test

import (
	"resource-plan-improvement/util"
	"testing"
)

func TestConvertStringArrToUintArr(t *testing.T) {
	strArr := []string{"1", "2", "3"}
	expectArr := []uint{1, 2, 3}
	actualArr, err := util.ConvertStringArrToUintArr(strArr)
	if err != nil {
		t.Fatalf(err.Error())
	}
	for i := 0; i < len(strArr); i++ {
		if expectArr[i] != actualArr[i] {
			t.Fatalf("expect: %d, actual: %d", expectArr[i], actualArr[i])
		}
	}
}
