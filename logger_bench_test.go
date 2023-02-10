package logger

import (
	"os"
	"testing"
)

func BenchmarkInfo(b *testing.B) {
	b.ResetTimer()
	f, _ := os.OpenFile("./test_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	config.Outputs = append(config.Outputs, f)
	for i := 0; i < b.N; i++ {
		Info("My Info Log Message")
	}
}
