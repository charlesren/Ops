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

// GetStCodeGood return status code
func GetStCodeGood(data float64, cfgItem string) (stcode int, good string) {
	StCode = 110
	vUpper := float64(999999999999)
	thordHolds := strings.Split(cfgItem, "#")[1]
	thordHold := strings.Split(thordHolds, ";")
	otherStCodeString := strings.Split(cfgItem, "#")[2]
	otherStCode, _ := strconv.Atoi(otherStCodeString)
	//get StCode
	for _, td := range thordHold {
		codeString := strings.Split(td, ",")[0]
		code, _ := strconv.Atoi(codeString)
		thordString := strings.Split(td, ",")[1]
		thord, _ := strconv.ParseFloat(thordString, 64)
		if thord >= data && thord < vUpper {
			StCode = code
			vUpper = thord
		}

	}
	if StCode == 110 {
		StCode = otherStCode
	}
	//get Good
	var Good string
	var before string
	for index, td := range thordHold {
		codeString := strings.Split(td, ",")[0]
		code, _ := strconv.Atoi(codeString)
		thordString := strings.Split(td, ",")[1]
		//thord, _ := strconv.ParseFloat(thordString, 64)
		if otherStCode == 0 {
			if index == (len(thordHold) - 1) {
				Good = "[" + thordString + " " + "Max]"
			}
		} else {
			if index == 0 && code == 0 {
				Good = "[0" + " " + thordString + "]"
				break
			} else {
				if code != 0 {
					before = thordString
				}
				if code == 0 {
					Good = "[" + before + " " + thordString + "]"
					break
				}
			}
		}
	}
	return StCode, Good
}
