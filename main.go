package main

import (
	"fmt"
	"log"
	"ops/src/entegor"
	"ops/src/filenum"
	"os"
	"path/filepath"
	"time"
)

func main() {
	err := os.Setenv("LANG", "en_US")
	if err != nil {
		log.Println(err)
	}
	walkDir := "/usr"
	filenum.ThordHold = 1000
	filepath.Walk(walkDir, filenum.CheckNum)
	fmt.Println(filenum.Files)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hostname)
	now := time.Now().Format(entegor.LongForm)
	fmt.Println(now)
	//	entegor.SaveData()
}
