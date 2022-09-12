package main

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()

func init() {
	// https://pkg.go.dev/time#pkg-constants
	// yyyy-MM-dd HH:mm:ss와 같은 구문 분석을 위하여 Mon Jan 2 15:04:05 MST 2006를 기준 시각으로 한다.
	log.SetFormatter(&prefixed.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	log.Info("Believe in yourself")
	log.Debug("No Silver bullet")
	log.Warn("I can eat glass and it doesn't hurt me.")
}
