package main

import (
	"bufio"
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
	var Message entegor.Message
	Message.WarnDesc = filenum.WarnDesc
	Message.Hostname = entegor.GetHostname()
	Message.CheckTime = time.Now().Format(entegor.LongForm)
	Message.ErrCode = filenum.ErrCode
	Message.Script = fullScriptName
	Message.GMESSENGER = entegor.GMESSENGER
	Message.HostIP = HostIP
	Message.WarnDesc = filenum.WarnDesc
	ini, err := os.Open(INIFile)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	iniScanner := bufio.NewScanner(ini)
	for iniScanner.Scan() {
		cfgItem := iniScanner.Text()
		walkDir := strings.Split(strings.Split(cfgItem, "|")[0], "=")[1]
		filenum.ThordHold, _ = strconv.Atoi(strings.Split(strings.Split(cfgItem, "#")[0], "|")[1])
		filepath.Walk(walkDir, filenum.CheckNum)
		Dirs := &filenum.Dirs
		var Data float64
		if *Dirs == nil {
			Data = float64(0)
		} else {
			Data = float64(1)
		}
		Message.StCode = entegor.GetStCode(Data, cfgItem)
		Message.OutDesc = walkDir
		Message.Threadhold = entegor.GetGood(cfgItem)
		Message.CheckData = strconv.FormatFloat(Data, 'f', -1, 64)
		var Msg string
		for _, file := range *Dirs {
			sysutil.AppendToFile(LogFile, file.Name+"   "+strconv.Itoa(file.Num)+"\n")
			Msg = Msg + file.Name + "   " + strconv.Itoa(file.Num) + ";"
		}
		Message.OutDesc = Message.OutDesc + "|" + Msg
		Message.OutHead = entegor.GetHead(cfgItem)
		Message.WarnDesc = Message.WarnDesc + ":" + Msg
		Message.WarnHead = entegor.GetWarningHead(cfgItem)
		entegor.SaveData(&Message, OutTmpFile, OutFile)
	}
}
