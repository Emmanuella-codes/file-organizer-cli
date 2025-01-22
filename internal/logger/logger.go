package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	logFile *os.File
}

func New(logDir string) (*Logger, error) {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	logPath := filepath.Join(logDir, fmt.Sprintf("file-organizer-%s.log", time.Now().Format("2006-01-02")))
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }

    return &Logger{logFile: logFile}, nil
}

func (l *Logger) Log(message string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_, err := fmt.Fprintf(l.logFile, "[%s] %s\n", timestamp, message)
	return err
}

func (l *Logger) Close() error {
	return l.logFile.Close()
}
