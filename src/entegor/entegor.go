package entegor

import (
	"fmt"
	"strconv"
	"strings"
)

// StCode is check result code
var StCode int

// SaveData save check result
func SaveData() string {
	cfgItem := "abc#0,25;1,30;2,35;4,40#5"
	thordHolds := strings.Split(cfgItem, "#")[1]
	fmt.Println(thordHolds)
	return thordHolds
}

// GetStCode return status code
func GetStCode(data float64, cfgItem string) int {
	StCode = 110
	vLower := float64(-99999999999)
	fmt.Printf("Default lower is %v\n", vLower)
	vUpper := float64(999999999999)
	fmt.Printf("Default upper is %v\n", vUpper)
	thordHolds := strings.Split(cfgItem, "#")[1]
	thordHold := strings.Split(thordHolds, ";")
	otherStCodeString := strings.Split(cfgItem, "#")[2]
	otherStCode, _ := strconv.Atoi(otherStCodeString)
	firstStCodeString := strings.Split(thordHold[0], ",")[0]
	firstStCode, _ := strconv.Atoi(firstStCodeString)
	var Thord float64
	var ThordZero string
	var ThordLast string
	var Right string
	for index, td := range thordHold {
		codeString := strings.Split(td, ",")[0]
		code, _ := strconv.Atoi(codeString)
		thordString := strings.Split(td, ",")[1]
		thord, _ := strconv.ParseFloat(thordString, 64)
		if thord >= data && thord < vUpper {
			StCode = code
			vUpper = thord
			Thord = thord
		}
		if code == 0 {
			ThordZero = thordString
		}
		if index == (len(thordHold) - 1) {
			ThordLast = thordString
		}
	}
	if StCode == 110 {
		StCode = otherStCode
	}
	if firstStCode < otherStCode {
		Right = "[0" + " " + ThordZero + "]"
	} else {
		Right = "[" + ThordLast + " " + "Max]"
	}
	fmt.Println(firstStCode)
	fmt.Println(Thord)
	fmt.Printf("ThordZero is %v\n", ThordZero)
	fmt.Printf("ThordLast is %v\n", ThordLast)
	fmt.Println(Right)
	return StCode
}
