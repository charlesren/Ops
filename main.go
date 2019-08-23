package main

import (
	"fmt"
	"ops/src/entegor"
	"ops/src/filenum"
	"os"
	"path/filepath"
)

func main() {
	walkDir := "/usr"
	filenum.ThordHold = 1000
	filepath.Walk(walkDir, filenum.CheckNum)
	fmt.Println(filenum.Files)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(hostname)
	}
	entegor.SaveData()
}
