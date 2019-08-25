package main

import (
	"bufio"
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

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
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
	inifile := "./inifile.ini"
	ini, err := os.Open(inifile)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	iniScanner := bufio.NewScanner(ini)
	for iniScanner.Scan() {
		line := iniScanner.Text()
		fmt.Println(line)
	}
	//	entegor.SaveData()
}
