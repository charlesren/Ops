package entegor

import (
	"fmt"
	"strconv"
	"strings"
)

// SaveData save check result
func SaveData() string {
	cfgItem := "abc#0,25;1,30;2,35;4,40#5"
	thordHolds := strings.Split(cfgItem, "#")[1]
	fmt.Println(thordHolds)
	return thordHolds
}

// GetStCode return status code
func GetStCode(data float64, cfgItem string) int {
	vLower := -99999999999
	vUpper := 999999999999
	thordHolds := strings.Split(cfgItem, "#")[1]
	thordHold := strings.Split(thordHolds, ";")
	for _, t := range thordHold {
		code := strings.Split(t, ",")[0]
		code1, _ := strconv.Atoi(code)
		fmt.Println(code1)
		thord := strings.Split(t, ",")[1]
		thord1, _ := strconv.ParseFloat(thord, 64)
		fmt.Println(thord1)
	}
	fmt.Println(thordHolds)
	fmt.Println(vLower)
	fmt.Println(vUpper)

	return 1
}
