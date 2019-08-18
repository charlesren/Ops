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
	StCode = 6
	vLower := float64(-99999999999)
	fmt.Printf("Default lower is %v\n", vLower)
	vUpper := float64(999999999999)
	fmt.Printf("Default upper is %v\n", vUpper)
	thordHolds := strings.Split(cfgItem, "#")[1]
	thordHold := strings.Split(thordHolds, ";")
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
	if StCode == 6 {
		fmt.Println("function err")
	}
	return StCode
}
