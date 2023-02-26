package validation

import (
	"fmt"
	"log"
	"os"
)

// Path to the log file
var LOGFILE = "../tmp/logger.log"

// Function that writes information on error to a passed *log.Logger object
func writeToLog(aLog *log.Logger, err error) {
	aLog.Println(err)
}

// Funcation that creates a *log.Logger object with specific flags.
// Flags used in creation: log.LstdFlags and log.Lshortfile
func logCreation() (*log.Logger, error) {
	f, err := os.OpenFile(LOGFILE, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("%w (%s)", ErrLogFileNotOpened, LOGFILE)

	}

	aLog := log.New(f, "ellofae-validator", log.LstdFlags)
	aLog.SetFlags(log.LstdFlags | log.Lshortfile)

	return aLog, nil
}

// Function for user's usage to clean up the log file from previous logss
func LogCleaner() {
	if err := os.Truncate(LOGFILE, 0); err != nil {
		log.Printf("Failed to trancate: %v\n", err)
	}
}
