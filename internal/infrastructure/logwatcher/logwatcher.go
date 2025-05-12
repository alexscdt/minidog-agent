package logwatcher

import (
	"log"
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

		output <- domain.LogEvent{
			FilePath:  path,
			Timestamp: time.Now(),
			Content:   line.Text,
		}
	}
}
