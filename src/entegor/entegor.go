package entegor

import (
	"strconv"
	"strings"
)

//LongForm define long time format
const LongForm = "2016-01-02 15:04:05"

//ShortForm define short time format
const ShortForm = "2016-01-02"

// StCode is check result code
var StCode int

// SaveData save check result
func SaveData(stCode int, cfgItem string, now string, data float64, descMsg string) string {
	var result string
	good := GetGood(cfgItem)
	stCodeString := strconv.Itoa(stCode)
	dataString := strconv.FormatFloat(data, 'E', 1, 64)
	if stCode == 0 {
		head := GetHead(cfgItem)
		result = head + "=" + stCodeString + "|" + now + "|" + dataString + "|" + good + "|" + descMsg
	} else {
		head := GetWarningHead(cfgItem)
		result = head + "=" + stCodeString + "|" + now + "|" + "AOMS" + "|" + dataString + "|" + good + "|" + descMsg
	}
	return result
}

//GetHead return head
func GetHead(cfgItem string) string {
	head := strings.Split(cfgItem, "=")[0]
	return head
}

//GetWarningHead return Warning head
func GetWarningHead(cfgItem string) string {
	tempHead := strings.SplitAfter(cfgItem, ",")
	warningHead := tempHead[0] + tempHead[1] + "-1"
	return warningHead
}

// GetStCode return status code
func GetStCode(data float64, cfgItem string) int {
	StCode = 110
	vUpper := float64(999999999999)
	thordHolds := strings.Split(cfgItem, "#")[1]
	thordHold := strings.Split(thordHolds, ";")
	otherStCodeString := strings.Split(cfgItem, "#")[2]
	otherStCode, _ := strconv.Atoi(otherStCodeString)
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
	return StCode
}

// GetGood return good range
func GetGood(cfgItem string) string {
	StCode = 110
	thordHolds := strings.Split(cfgItem, "#")[1]
	thordHold := strings.Split(thordHolds, ";")
	otherStCodeString := strings.Split(cfgItem, "#")[2]
	otherStCode, _ := strconv.Atoi(otherStCodeString)
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
				} else {
					Good = "[" + before + " " + thordString + "]"
					break
				}
			}
		}
	}
	return Good
}
