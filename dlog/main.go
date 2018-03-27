package dlog

import (
	"log"
	"os"
)

var Stdlog, Errlog *log.Logger

func init() {
	Stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	Errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}
