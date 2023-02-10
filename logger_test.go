package logger

import (
	"os"
	"testing"
)

func TestLogToFile(t *testing.T) {
	f, _ := os.OpenFile("./test_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	config.Outputs = append(config.Outputs, f)
	Info("Test Log To File", "Oke")
}
