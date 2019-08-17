package main

import (
	"fmt"
	"ops/src/entegor"
	"ops/src/filenum"
	"path/filepath"
)

func main() {
	walkDir := "/usr"
	filenum.ThordHold = 1000
	filepath.Walk(walkDir, filenum.CheckNum)
	fmt.Println(filenum.Files)
	entegor.SaveData()
}
