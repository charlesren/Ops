package filenum

import (
	"fmt"
	"ops/src/entegor"
	"path/filepath"
	"testing"
)

func main() {
	WalkDir := "/usr"
	ThordHold = 1000
	filepath.Walk(WalkDir, CheckNum)
	fmt.Println(Files)
	entegor.SaveData()
}
func TestCheckNum(t *testing.T) {
	WalkDir := "/usr"
	ThordHold = 1000
	filepath.Walk(WalkDir, CheckNum)
	fmt.Println(Files)
	entegor.SaveData()
}
