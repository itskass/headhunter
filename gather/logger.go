package gather

import (
	"log"
	"os"
)

var (
	Log     = log.New(os.Stdout, "[hh]", log.Ldate|log.Ltime)
	SyncLog = log.New(os.Stdout, "[hh]", log.Ldate|log.Ltime)
)
