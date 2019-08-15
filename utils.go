package ops

import (
	"fmt"
	"strings"
)

// SaveData save check result
func SaveData() string {
	cfgItem := "abc#1,25;1,30;2,35;4,40#5"
	checkHold := strings.Split(cfgItem, "#")[1]
	fmt.Println(checkHold)
	return checkHold
}
