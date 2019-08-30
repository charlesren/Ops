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
	entegor.SetLang()
	INIFile := os.Args[1]
	HostIP12 := os.Args[2]
	HostIP := os.Args[3]
	_, fullScriptName := filepath.Split(os.Args[0])
	scriptName := strings.Split(fullScriptName, ".")[0]
	LogFile, OutTmpFile, OutFile := entegor.PrepareFile(HostIP12, scriptName)
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
		DataString := strconv.FormatFloat(Data, 'f', -1, 64)
		var WarnMsg string
		fmt.Println(filenum.Files)
		for _, file := range filenum.Files {
			sysutil.AppendToFile(LogFile, file.Name+"   "+strconv.Itoa(file.Num)+"\n")
			WarnMsg = WarnMsg + file.Name + "   " + strconv.Itoa(file.Num) + "\n"
		}
		var result string
		head := entegor.GetHead(cfgItem)
		result = head + "=" + stCodeString + "|" + checkTime + "|" + DataString + "|" + good + "|" + descMsg + "\n"
		sysutil.AppendToFile(OutTmpFile, result)
		sysutil.AppendToFile(OutFile, result)
		if stCode != 0 {
			head := entegor.GetWarningHead(cfgItem)
			result = head + "=" + stCodeString + "|" + checkTime + "|" + "AOMS" + "|" + fullScriptName + "|" + filenum.ErrCode + "|" + hostname + "|" + HostIP + "|" + "" + "|" + "" + "|" + WarnMsg + "\n"
			sysutil.AppendToFile(OutTmpFile, result)
			sysutil.AppendToFile(OutFile, result)
		}
	}
	//sysutil.AppendToFile(OutFile, result)
}
