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

func TestGetStCode(t *testing.T) {
	cfgItem := "abc=cpu#0,25;1,30;2,35;4,40#5"
	data := float64(37)
	stcode := GetStCode(data, cfgItem)
	if stcode != 4 {
		t.Errorf(`stcode is %v ,should be 4!!!`, stcode)
	}
	cfgItem1 := "abc=cpu#5,0;4,10;3,20;2,30;1,40#0"
	data1 := float64(26.5)
	stcode1 := GetStCode(data1, cfgItem1)
	if stcode1 != 2 {
		t.Errorf(`stcode1 is %v ,should be 2!!!`, stcode1)
	}
	cfgItem2 := "abc=cpu#4,10;2,30;0,40#5"
	data2 := float64(36.5)
	stcode2 := GetStCode(data2, cfgItem2)
	if stcode2 != 0 {
		t.Errorf(`stcode2 is %v ,should be 0!!!`, stcode2)
	}
}
func TestGetGood(t *testing.T) {
	cfgItem := "abc=cpu#0,25;1,30;2,35;4,40#5"
	good := GetGood(cfgItem)
	if good != "[0 25]" {
		t.Errorf(`good should be [0 25],not %v`, good)
	}
	cfgItem1 := "abc=cpu#5,0;4,10;3,20;2,30;1,40#0"
	good1 := GetGood(cfgItem1)
	if good1 != "[40 Max]" {
		t.Errorf(`good1 should be [40 Max],not %v`, good1)
	}
	cfgItem2 := "abc=cpu#4,10;2,30;0,40#5"
	good2 := GetGood(cfgItem2)
	if good2 != "[30 40]" {
		t.Errorf(`good2 should be [30 40],not %v`, good2)
	}
}
func TestGetHead(t *testing.T) {
	cfgItem := "abc,000,111=cpu#0,25;1,30;2,35;4,40#5"
	head := GetHead(cfgItem)
	if head != "abc,000,111" {
		t.Errorf(`head should be "abc,000,111",not %v`, head)
	}
}
func TestGetWarningHead(t *testing.T) {
	cfgItem := "abc,000,111=cpu#0,25;1,30;2,35;4,40#5"
	head := GetWarningHead(cfgItem)
	if head != "abc,000,-1" {
		t.Errorf(`head should be "abc,000,-1",not %v`, head)
	}
}
