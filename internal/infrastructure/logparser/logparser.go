package logparser

import (
	"encoding/json"
	"github.com/alexscdt/minidog-agent/internal/domain"
	"strings"
	"time"
)

type DockerRawLog struct {
	Log    string `json:"log"`
	Stream string `json:"stream"`
	Time   string `json:"time"`
}

func DetectLogLevel(line string) string {
	line = strings.ToLower(line)
	switch {
	case strings.Contains(line, "error"):
		return "error"
	case strings.Contains(line, "warn"):
		return "warn"
	case strings.Contains(line, "debug"):
		return "debug"
	case strings.Contains(line, "fatal"):
		return "fatal"
	default:
		return "info"
	}
}

func ParseDockerLogLine(path string, line string) (domain.LogEvent, error) {
	var raw DockerRawLog
	err := json.Unmarshal([]byte(line), &raw)
	if err != nil {
		return domain.LogEvent{}, err
	}

	timestamp, err := time.Parse(time.RFC3339Nano, raw.Time)
	if err != nil {
		timestamp = time.Now()
	}

	return domain.LogEvent{
		FilePath:  path,
		Timestamp: timestamp,
		Content:   strings.TrimSpace(raw.Log),
		Level:     DetectLogLevel(raw.Log),
	}, nil
}
