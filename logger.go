package logger

import (
	"bytes"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var (
	FLAG_DEBUGP = 5
	FLAG_DEBUG  = 4
	FLAG_INFO   = 3
	FLAG_WARN   = 2
	FLAG_ERROR  = 1

	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

type LoggerConfig struct {
	Flag            int
	Outputs         []*os.File
	MessageChanSize int
}

type Logger struct {
	Config      *LoggerConfig
	messageChan chan []byte
	exitChan    chan bool
	running     int32
}

var config = &LoggerConfig{
	Flag:            FLAG_DEBUGP,
	Outputs:         []*os.File{os.Stdout},
	MessageChanSize: 1000,
}

var logger = &Logger{
	Config:   config,
	exitChan: make(chan bool),
	running:  0,
}

func RunLogger() {
	logger.messageChan = make(chan []byte, logger.Config.MessageChanSize)
	atomic.AddInt32(&logger.running, 1)
	go logger.runWriteAsyncWriteToOutputs()
}

func StopLogger() {
	logger.exitChan <- true
}

func checkRun() {
	running := atomic.LoadInt32(&logger.running)
	if running == 0 {
		RunLogger()
	}
}

func SetConfig(newConfig *LoggerConfig) {
	config = newConfig
}

func DebugP(message interface{}, a ...interface{}) {
	if config.Flag < FLAG_DEBUGP {
		return
	}
	checkRun()
	logger.messageChan <- getLogBuffer(Purple, "DEBUG_P", message, a)
}

func Debug(message interface{}, a ...interface{}) {
	if config.Flag < FLAG_DEBUG {
		return
	}
	checkRun()
	logger.messageChan <- getLogBuffer(Cyan, "DEBUG", message, a)
}

func Info(message interface{}, a ...interface{}) {
	if config.Flag < FLAG_INFO {
		return
	}
	checkRun()
	logger.messageChan <- getLogBuffer(Green, "INFO", message, a)
}

func Warn(message interface{}, a ...interface{}) {
	if config.Flag < FLAG_WARN {
		return
	}
	checkRun()
	logger.messageChan <- getLogBuffer(Yellow, "WARN", message, a)
}

func Error(message interface{}, a ...interface{}) {
	if config.Flag < FLAG_ERROR {
		return
	}
	checkRun()
	logger.messageChan <- getLogBuffer(Red, "ERROR", message, a)
}

func getLogBuffer(color string, prefix string, message interface{}, a []interface{}) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(color)
	buffer.WriteString("[")
	buffer.WriteString(prefix)
	buffer.WriteString("][")
	buffer.WriteString(time.Now().Format(time.RFC822Z))
	buffer.WriteString("] ")
	buffer.WriteString(fmt.Sprintf("%v ", message))
	for _, v := range a {
		buffer.WriteString(fmt.Sprintf("\n%v", v))
	}
	buffer.WriteString(Reset)
	buffer.WriteString("\n")
	return buffer.Bytes()
}

func (logger *Logger) runWriteAsyncWriteToOutputs() {
	for {
		select {
		case message := <-logger.messageChan:
			for i := range config.Outputs {
				config.Outputs[i].Write(message)
			}
		case <-logger.exitChan:
			return
		}
	}
}
