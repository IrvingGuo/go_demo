package util

import "strconv"

func ConvertStringArrToUintArr(strArr []string) (uintArr []uint, err error) {
	var intItem int
	for _, str := range strArr {
		if intItem, err = strconv.Atoi(str); err != nil {
			return
		} else {
			uintArr = append(uintArr, uint(intItem))
		}
	}
	return
}
