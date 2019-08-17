package entegor

import (
	"fmt"
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
	fmt.Println(thordHolds)
	fmt.Println(vLower)
	fmt.Println(vUpper)
	return 1
}
