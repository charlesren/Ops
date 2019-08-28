package main

import (
	"bufio"
	"fmt"
	"log"
	"ops/src/entegor"
	"ops/src/filenum"
	"ops/src/sysutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := os.Setenv("LANG", "en_US")
	if err != nil {
		log.Println(err)
	}
	INIFile := os.Args[1]
	HostIP12 := os.Args[2]
	HostIP := os.Args[3]
	/*
		INIFile := "./inifile.ini"
		HostIP12 := "011111111111"
		HostIP := "11.111.111.111"
	*/
	fmt.Println(INIFile, HostIP12, HostIP)
	WorkDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_, fullScriptName := filepath.Split(os.Args[0])
	scriptName := strings.Split(fullScriptName, ".")[0]
	TmpDir := filepath.Join(WorkDir, "temp")
	OutDir := filepath.Join(WorkDir, "out")
	LogDir := filepath.Join(WorkDir, "log")
	Dirs := []string{TmpDir, OutDir, LogDir}
	for _, dir := range Dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}
	}
	LogFileName := scriptName + HostIP12 + ".log"
	OutTmpFileName := scriptName + HostIP12 + ".out"
	OutFileName := "check" + HostIP12 + ".out"
	LogFile := filepath.Join(LogDir, LogFileName)
	OutTmpFile := filepath.Join(TmpDir, OutTmpFileName)
	OutFile := filepath.Join(OutDir, OutFileName)
	sysutil.WriteToFile(OutTmpFile, "")
	sysutil.WriteToFile(LogFile, "")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hostname)
	checkTime := time.Now().Format(entegor.LongForm)
	ini, err := os.Open(INIFile)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	iniScanner := bufio.NewScanner(ini)
	for iniScanner.Scan() {
		cfgItem := iniScanner.Text()
		walkDir := strings.Split(strings.Split(cfgItem, "|")[0], "=")[1]
		filenum.ThordHold, _ = strconv.Atoi(strings.Split(strings.Split(cfgItem, "#")[0], "|")[1])
		fmt.Println(filenum.ThordHold)
		filepath.Walk(walkDir, filenum.CheckNum)
		fmt.Println(filenum.Files)
		var Data float64
		if filenum.Files == nil {
			Data = float64(0)
		} else {
			Data = float64(1)
		}
		stCode := entegor.GetStCode(Data, cfgItem)
		descMsg := walkDir
		good := entegor.GetGood(cfgItem)
		stCodeString := strconv.Itoa(stCode)
		DataString := strconv.FormatFloat(Data, 'E', 1, 64)
		var WarnMsg string
		fmt.Println(filenum.Files)
		for _, file := range filenum.Files {
			sysutil.AppendToFile(LogFile, file.Name+"   "+strconv.Itoa(file.Num)+"\n")
			WarnMsg = WarnMsg + file.Name + "   " + strconv.Itoa(file.Num) + "\n"
		}
		var result string
		if stCode == 0 {
			head := entegor.GetHead(cfgItem)
			result = head + "=" + stCodeString + "|" + checkTime + "|" + DataString + "|" + good + "|" + descMsg + "\n"
		} else {
			head := entegor.GetWarningHead(cfgItem)
			result = head + "=" + stCodeString + "|" + checkTime + "|" + "AOMS" + "|" + fullScriptName + "|" + filenum.ErrCode + "|" + hostname + "|" + HostIP + "|" + "" + "|" + "" + "|" + WarnMsg + "\n"
		}
		sysutil.AppendToFile(OutTmpFile, result)
		sysutil.AppendToFile(OutFile, result)
	}
	//sysutil.AppendToFile(OutFile, result)
}
