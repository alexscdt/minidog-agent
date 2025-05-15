package logwatcher

import (
	"log"
	"strings"
	"time"

	"github.com/alexscdt/minidog-agent/internal/domain"
	"github.com/hpcloud/tail"
)

func WatchFile(path string, output chan<- domain.LogEvent) {
	t, err := tail.TailFile(path, tail.Config{
		Follow:    true,
		ReOpen:    true,
		MustExist: true,
		Poll:      true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	})
	if err != nil {
		log.Printf("Erreur d'ouverture du fichier %s : %v", path, err)
		return
	}

	for line := range t.Lines {
		if line.Err != nil {
			log.Printf("Erreur de lecture sur %s : %v", path, line.Err)
			continue
		}

		level := "info"
		if strings.Contains(strings.ToLower(line.Text), "error") {
			level = "error"
		}

		output <- domain.LogEvent{
			FilePath:  path,
			Timestamp: time.Now(),
			Content:   line.Text,
			Level:     level,
		}
	}
}
