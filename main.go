package main

import (
	"bufio"
	"fmt"
	"log"
	"ops/src/entegor"
	"ops/src/filenum"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	err := os.Setenv("LANG", "en_US")
	if err != nil {
		log.Println(err)
	}
	//INIFile := os.Args[1]
	//HostIP12 := os.Args[2]
	//HostIP := os.Args[3]
	INIFile := "./inifile.ini"
	HostIP12 := "011111111111"
	HostIP := "11.111.111.111"
	fmt.Println(INIFile, HostIP12, HostIP)
	WorkDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(WorkDir)
	_, fullScriptName := filepath.Split(os.Args[0])
	scriptName := strings.Split(fullScriptName, ".")[0]
	TmpDir := filepath.Join(WorkDir, "temp")
	OutDir := filepath.Join(WorkDir, "out")
	LogDir := filepath.Join(WorkDir, "log")
	LogFileName := scriptName + HostIP12 + ".log"
	OutTmpFileName := scriptName + HostIP12 + ".out"
	OutFileName := "check" + HostIP12 + ".out"
	LogFile := filepath.Join(LogDir, LogFileName)
	OutTmpFile := filepath.Join(TmpDir, OutTmpFileName)
	OutFile := filepath.Join(OutDir, OutFileName)
	fmt.Println(LogFile, OutTmpFile, OutFile)
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
	ini, err := os.Open(INIFile)
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
