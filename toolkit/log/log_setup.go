package log

import (
	"fmt"
	"io"
	stdLog "log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/log"
)

// NewDevLogger logger.
// Pretty logging for development mode.
// Not recommended for production use.
// If static fields are provided those values will define
// the default static fields for each new built instance
// if they were not yet configured.
func NewDevLogger(batchCfg *BatchConfig, stfields ...interface{}) *Logger {
	var writer io.Writer

	writer = os.Stdout

	if batchCfg != nil && batchCfg.Enabled {
		writer = diode.NewWriter(writer, batchCfg.MaxLines, batchCfg.Interval, func(missed int) {
			stdLog.Printf("Logger Dropped %d messages", missed)
		})
	}

	output := zerolog.ConsoleWriter{Out: writer, TimeFormat: time.RFC3339}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("** %s **", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s=", i)
	}
	output.FormatErrFieldValue = func(i interface{}) string {
		if e, ok := i.(error); ok {
			return e.Error()
		}

		return fmt.Sprintf("%s", i)
	}
	output.FormatTimestamp = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	output.FormatCaller = func(i interface{}) string {
		var sb strings.Builder

		callers := strings.Split(fmt.Sprintf("%s", i), "/")

		for i, v := range callers {
			if i == len(callers)-1 {
				sb.WriteString(callers[i])
				return sb.String()
			}

			if i == len(callers)-2 {
				sb.WriteString(fmt.Sprintf("%s/", callers[i]))
				continue
			}

			if i == 0 && v == "vendor" {
				sb.WriteString("vendor/")
				continue
			}

			chars := []rune(v)
			if len(chars) > 0 {
				sb.WriteRune(chars[0])
				sb.WriteString("/")
			}
		}

		return sb.String()
	}
	output.FormatErrFieldName = func(i interface{}) string {
		return "error="
	}

	stdl := zerolog.New(output).With().
		Timestamp().
		CallerWithSkipFrameCount(constant.LogSkipCallerCount).
		Logger()
	errl := zerolog.New(output).With().
		Timestamp().
		CallerWithSkipFrameCount(constant.LogSkipCallerCount).
		Logger()

	// level := LevelDebug

	setLogLevel(&stdl, constant.LogLevelDebug)
	setLogLevel(&errl, constant.LogLevelDebug)

	l := &Logger{
		Level:  constant.LogLevelDebug,
		StdLog: stdl,
		ErrLog: errl,
		logFmt: true,
	}

	// if len(stfields) > 1 && !cfg.configured {
	if !cfg.configured {
		setup(constant.LogLevelDebug, true, batchCfg, stfields)

		defaultLogger = l
	}

	return l
}

// NewLogger logger.
// If static fields are provided those values will define
// the default static fields for each new built instance
// if they were not yet configured.
func NewLogger(level Level, batchCfg *BatchConfig, stfields ...interface{}) *Logger {
	var (
		stdWriter io.Writer
		errWriter io.Writer
	)

	stdWriter = os.Stdout
	errWriter = os.Stderr

	if batchCfg != nil && batchCfg.Enabled {
		stdWriter = diode.NewWriter(stdWriter, batchCfg.MaxLines, batchCfg.Interval, func(missed int) {
			stdLog.Printf("Logger Dropped %d messages\n", missed)
		})
		errWriter = diode.NewWriter(errWriter, batchCfg.MaxLines, batchCfg.Interval, func(missed int) {
			stdLog.Printf("Logger Dropped %d messages\n", missed)
		})
	}

	stdl := log.Output(stdWriter).With().
		Timestamp().
		CallerWithSkipFrameCount(constant.LogSkipCallerCount).
		Logger()
	errl := log.Output(errWriter).With().
		Timestamp().
		CallerWithSkipFrameCount(constant.LogSkipCallerCount).
		Logger()

	setLogLevel(&stdl, level)
	setLogLevel(&errl, level)

	l := &Logger{
		Level:  level,
		StdLog: stdl,
		ErrLog: errl,
	}

	if len(stfields) > 1 && !cfg.configured {
		setup(level, false, batchCfg, stfields)

		defaultLogger = l
	}

	return l
}

// Set the base package logger.
func Set(l *Logger) {
	defaultLogger = l
}

// Set default package logger.
// Can be used chained with NewLogger to create a new one,
// set it up as package default logger and get it for use in one step.
// i.e:
// logger := log.NewLogger(log.Debug, "name", "version", "revision").Set().
func (l *Logger) Set() *Logger {
	defaultLogger = l
	return defaultLogger
}

// NewFromConfig returns logger based on config file
//
// log:
//
//	level: info
//	json-enabled: false
//	batch:
//	  enabled: false
//	  max-lines: 1000
//	  interval: 15ms
func NewFromConfig() (l *Logger, err error) {
	logJSONFormatStr := os.Getenv("LOG_JSON_ENABLED")
	logLevel := os.Getenv("LOG_LEVEL")
	enabledStr := os.Getenv("LOG_BATCH_ENABLED")
	intervalStr := os.Getenv("LOG_BATCH_INTERVAL")
	maxLinesStr := os.Getenv("LOG_BATCH_MAX_LINES")

	logJSONFormat, err := strconv.ParseBool(logJSONFormatStr)
	if err != nil {
		return nil, errors.Wrapf(err, "error parse bool on enable json log env : %s", logJSONFormatStr)
	}

	enabled, err := strconv.ParseBool(enabledStr)
	if err != nil {
		return nil, errors.Wrapf(err, "error parse bool on enable batch log env : %s", enabledStr)
	}

	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		return nil, errors.Wrapf(err, "error parse duration on interval batch log env : %s", intervalStr)
	}

	maxLines, err := strconv.Atoi(os.Getenv("LOG_BATCH_MAX_LINES"))
	if err != nil {
		return nil, errors.Wrapf(err, "error parse int on max lines batch log env : %s", maxLinesStr)
	}

	logBatchCfg := &BatchConfig{
		Enabled:  enabled,
		Interval: interval,
		MaxLines: maxLines,
	}

	if logJSONFormat {
		// Use JSONFormatter logger for non development environment
		l = NewLogger(GetLevelFromString(logLevel), logBatchCfg)
	} else {
		l = NewDevLogger(logBatchCfg)
	}

	return l, nil
}
