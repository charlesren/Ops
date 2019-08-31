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
		fmt.Println(filenum.ThordHold)
		filepath.Walk(walkDir, filenum.CheckNum)
		fmt.Println(filenum.Files)
		var Data float64
		if filenum.Files == nil {
			Data = float64(0)
		} else {
			Data = float64(1)
		}
		Message.StCode = entegor.GetStCode(Data, cfgItem)
		Message.OutDesc = walkDir
		Message.Threadhold = entegor.GetGood(cfgItem)
		//	stCodeString := strconv.Itoa(Message.StCode)
		Message.CheckData = strconv.FormatFloat(Data, 'f', -1, 64)
		var Msg string
		fmt.Println(filenum.Files)
		for _, file := range filenum.Files {
			sysutil.AppendToFile(LogFile, file.Name+"   "+strconv.Itoa(file.Num)+"\n")
			Msg = Msg + file.Name + "   " + strconv.Itoa(file.Num) + ";"
		}
		Message.OutDesc = Message.OutDesc + "|" + Msg
		Message.OutHead = entegor.GetHead(cfgItem)
		outData := fmt.Sprintf("%v=%v|%v|%v|%v|%v\n", Message.OutHead, Message.StCode, Message.CheckTime, Message.CheckData, Message.Threadhold, Message.OutDesc)
		sysutil.AppendToFile(OutTmpFile, outData)
		sysutil.AppendToFile(OutFile, outData)
		if Message.StCode != 0 {
			Message.WarnDesc = Message.WarnDesc + ":" + Msg
			Message.WarnHead = entegor.GetWarningHead(cfgItem)
			warnData := fmt.Sprintf("%v=%v|%v|%v|%v|%v|%v|%v|%v|%v|%v\n", Message.WarnHead, Message.StCode, Message.CheckTime, Message.GMESSENGER, Message.Script, Message.ErrCode, Message.Hostname, Message.HostIP, Message.CheckData, Message.Threadhold, Message.WarnDesc)
			sysutil.AppendToFile(OutTmpFile, warnData)
			sysutil.AppendToFile(OutFile, warnData)
		}
	}
	//sysutil.AppendToFile(OutFile, result)
}
