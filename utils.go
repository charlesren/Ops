package ops

import (
	"fmt"
	"strings"
)

// SaveData save check result
func SaveData() string {
	cfgItem := "abc#1,25;1,30;2,35;4,40#5"
	thordHolds := strings.Split(cfgItem, "#")[1]
	fmt.Println(thordHolds)
	return thordHolds
}

// GetStCode return status code
func GetStCode(data float64, cfgItem string) int {
	thordHolds := strings.Split(cfgItem, "#")[1]
	fmt.Println(thordHolds)
	return 1
}
