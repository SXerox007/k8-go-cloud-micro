package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	isTerminal  atomic.Value
	isTest      bool
	testingInit sync.Once
)

func init() {
	isTerminal.Store(terminal.IsTerminal(int(os.Stdout.Fd())))
	isTest = strings.HasSuffix(os.Args[0], ".test")
}

// NewLogger creates a new logrus logger
func NewLogger() *logrus.Logger {
	// Prepare a new logger
	logger := logrus.New()

	logger.Level = logrus.InfoLevel
	if isTest {
		testingInit.Do(func() { testing.Init() })

		if !testing.Verbose() {
			// Keep the tests quiet, unless -test.v is used.
			logger.Level = logrus.FatalLevel
		}
	}

	logger.Out = os.Stdout
	logger.SetReportCaller(true)

	formatter := &logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		DisableColors:   false,
	}
	logger.Formatter = formatter

	return logger
}

func WithError(err error, logger logrus.FieldLogger) *logrus.Entry {
	return logger.WithError(err).WithField("stacktrace", fmt.Sprintf("%+v", err))
}

func WithEncryptedField(k string, v interface{}, logger logrus.FieldLogger) *logrus.Entry {
	return logger.WithField(k, v)
}
