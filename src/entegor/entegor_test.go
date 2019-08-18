package entegor

import (
	"testing"
)

// SaveData save check result
/*
func TestSaveData(t *testing.T) {
	cfgItem := "abc#0,25;1,30;2,35;4,40#5"
	thordHolds := strings.Split(cfgItem, "#")[1]
	fmt.Println(thordHolds)
}
*/

// GetStCode return status code
func TestGetStCode(t *testing.T) {
	cfgItem := "abc#0,25;1,30;2,35;4,40#5"
	data := 37.7
	stcode := GetStCode(data, cfgItem)
	if stcode != 4 {
		t.Errorf(`stcode is %v ,should be 4!!!`, stcode)
	}
}
