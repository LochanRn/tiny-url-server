package logger

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	logg           *log.Logger
	prefixedLogger *log.Logger
)

type LogPrefix string

const (
	LogPrefixConfiguration LogPrefix = "[configuration]"
	LogPrefixSpec          LogPrefix = "[spec]"
	LogPrefixConnectivity  LogPrefix = "[connectivity]"
	LogPrefixActions       LogPrefix = "[actions]"
)

const (
	SiteInstallerConfigLogDir = "logs"
	TinyUrlLog                = "tiny-server.log"
)

func InitLogger(logLevel string) {
	logg = log.New()
	logg.SetReportCaller(true)
	logg.SetFormatter(&TinyUrlFormatter{})
	if logLevel == "debug" {
		logg.SetLevel(log.DebugLevel)
	}
	logg.SetOutput(os.Stdout)
	InitPrefixLogger(filepath.Join(SiteInstallerConfigLogDir, TinyUrlLog))
}

func Debug(format string, args ...interface{}) {
	entry := createEntry()
	entry.Debug(fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	entry := createEntry()
	entry.Info(fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	entry := createEntry()
	entry.Error(fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	entry := createEntry()
	entry.Warn(fmt.Sprintf(format, args...))
}

type TinyUrlFormatter struct {
}

func (f *TinyUrlFormatter) Format(entry *log.Entry) ([]byte, error) {
	return logrusFormat(entry, getLevelColor(entry.Level))
}

func logrusFormat(entry *log.Entry, levelColor int) ([]byte, error) {
	b := &bytes.Buffer{}
	level := strings.ToUpper(entry.Level.String())

	enableColor := true
	// Write timestamp
	if levelColor >= 0 {
		b.WriteString("\x1b[0m")
	} else {
		enableColor = false
	}
	b.WriteString(fmt.Sprintf("%s ", entry.Time.Format(time.StampMilli)))

	// Write level
	b.WriteString(colorPrintf(enableColor, levelColor, "%s", level))

	// Write callers
	if len(entry.Data) != 0 {
		fmt.Fprintf(b, "%s", getCallerString(entry.Data))
	}

	// Write message
	b.WriteString(colorPrintf(enableColor, 0, "[%s] ", entry.Message))

	// Write fields
	if len(entry.Data) != 0 {
		fields := make([]string, 0, len(entry.Data))
		for field := range entry.Data {
			fields = append(fields, field)
		}

		sort.Strings(fields)
		b.WriteByte('[')
		for _, field := range fields {
			fmt.Fprintf(b, "%s", colorPrintf(enableColor, levelColor, "%s=", field))
			fmt.Fprintf(b, "%s", colorPrintf(enableColor, 0, "%v ", entry.Data[field]))
		}
		b.WriteByte('\b')
		b.WriteByte(']')
	}
	b.WriteByte('\n')

	return b.Bytes(), nil
}

func getCallerString(data log.Fields) string {
	b := strings.Builder{}
	if len(data) != 0 {
		b.WriteString(fmt.Sprintf(" [%v/%v] ", data["file"], data["func"]))
		delete(data, "func")
		delete(data, "file")
	}
	return b.String()
}

func colorPrintf(enableColor bool, colorLevel int, format string, a ...interface{}) string {
	args := make([]interface{}, 0)
	if enableColor {
		format = "\x1b[%dm" + format
		if a != nil {
			args = append(args, colorLevel)
		}
	}
	args = append(args, a...)
	return fmt.Sprintf(format, args...)
}

func createEntry() *log.Entry {
	entry := log.NewEntry(logg)
	entry = entry.WithFields(getReportCallerFields())
	return entry
}

var (
	minCallerSkip = 3
	pkgCallerSkip = 4
	infoLevel     = false
	reportCaller  = true
)

func getReportCallerFields() log.Fields {
	var fun string
	var skip = minCallerSkip
	if infoLevel {
		skip++
		infoLevel = false
	}
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "unknown"
		line = 0
		fun = "unknown"
	}

	if fun != "unknown" {
		fun = runtime.FuncForPC(pc).Name()
	}

	return log.Fields{
		"file": fmt.Sprintf("%s:%d", trimFileName(file), line),
		"func": trimFileName(fun),
	}
}

func trimFileName(str string) string {
	i := strings.LastIndex(str, string(os.PathSeparator))
	if i > 0 && i < len(str) {
		return str[i+1:]
	}
	return str
}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func getLevelColor(level log.Level) int {
	switch level {
	case log.DebugLevel, log.TraceLevel:
		return gray
	case log.WarnLevel:
		return yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		return red
	default:
		return blue
	}
}

func PrintError(format string, args ...interface{}) {
	Error(format, args...)
}

func IsFileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func InitPrefixLogger(path string) {
	dirPath := filepath.Dir(path)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		Error("error creating prefixed logger file directory %s . Failed with error: %v", dirPath, err)
		return
	}
	if !IsFileExists(path) {
		_, err := os.Create(path)
		if err != nil {
			Error("error creating prefixed logger file %s . Failed with error: %v", path, err)
			return
		}
		Info("Successfully created %s", path)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		Error("error opening prefixed logger file %s . Failed with error: %v", path, err)
		return
	}
	logger := log.New()
	logger.SetOutput(file)
	prefixedLogger = logger
	Info("Successfully initialised PrefixedLogger")
}

// LogErrorWithPrefix logs errors encountered to stdOut and local-mgmt.log for any CRUD operation
func LogErrorWithPrefix(prefix LogPrefix, operation string, err error) {
	err = fmt.Errorf("%s Operation '%s' failed. Reason %s", prefix, operation, err.Error())
	if prefixedLogger != nil {
		prefixedLogger.Error(err)
	} else {
		Error("PrefixedLogger not initialised")
	}
	Error(err.Error()) // log in stdout
}

func LogWithPrefix(prefix LogPrefix, operation, message string) {
	msg := fmt.Sprintf("%s Operation %s performed successfully.", prefix, operation)
	if len(message) > 0 {
		msg += message
	}
	if prefixedLogger != nil {
		prefixedLogger.Info(msg)
	} else {
		Error("PrefixedLogger not initialised")
	}
	Info(msg) // log in stdout
}
