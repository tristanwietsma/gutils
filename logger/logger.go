package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	// CRITICAL log level
	CRITICAL *log.Logger

	// ERROR log level
	ERROR *log.Logger

	// WARNING log level
	WARNING *log.Logger

	// INFO log level
	INFO *log.Logger
)

// Init sets up loggers.
func Init() error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	CRITICAL = log.New(
		os.Stdout,
		fmt.Sprintf("HOST=%s LEVEL=CRITICAL ", hostname),
		log.LstdFlags,
	)

	ERROR = log.New(
		os.Stdout,
		fmt.Sprintf("HOST=%s LEVEL=ERROR ", hostname),
		log.LstdFlags,
	)

	WARNING = log.New(
		os.Stdout,
		fmt.Sprintf("HOST=%s LEVEL=WARNING ", hostname),
		log.LstdFlags,
	)

	INFO = log.New(
		os.Stdout,
		fmt.Sprintf("HOST=%s LEVEL=INFO ", hostname),
		log.LstdFlags,
	)

	return nil
}
