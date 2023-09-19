package src

import (
	"os"
	"time"
)

// Aquí se hace toda lo que tenga que ver con el log

type AppLog struct {
	File    string
	Line    uint
	Message string
	LogDate time.Time
}

func (l *AppLog) FreateOrOpen(file string) {
	_, err := os.Stat(file)
	if err != nil {
		// el archivo no existe.
	}

}

func (l *AppLog) update() {

}
