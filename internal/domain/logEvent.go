package domain

import "time"

type LogEvent struct {
	FilePath  string    `json:"file_path"`
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
	Level     string    `json:"level"`
}
