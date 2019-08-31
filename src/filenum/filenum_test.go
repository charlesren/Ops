package filenum

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestCheckNum(t *testing.T) {
	WalkDir := "/usr"
	ThordHold = 1000
	filepath.Walk(WalkDir, CheckNum)
	fmt.Println(Dirs)
}
