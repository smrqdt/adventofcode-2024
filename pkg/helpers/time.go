package helpers

import (
	"time"

	"github.com/charmbracelet/log"
)

func TrackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Warnf("%s took %s\n", name, elapsed)
}
