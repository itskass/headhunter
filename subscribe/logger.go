package subscribe

import (
	"log"
	"os"
)

var (
	Log = log.New(os.Stdout, "[hh]", log.Ldate|log.Ltime)
)
