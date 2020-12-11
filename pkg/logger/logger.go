package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var (
	log      *logrus.Logger
	repoPath string
)

// Initialization 初始化
func Initialization(logPath string) *logrus.Logger {
	repoPath = path.Join(strings.Replace(os.Getenv("GOPATH"), "\\", "/", -1), "src/github.com/zi-ao")

	// 日志文件
	fileName := path.Join(logPath, "api")
	// 写入文件
	writer, err := os.OpenFile(fileName+".log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	log = logrus.New()
	log.SetOutput(writer)
	log.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y-%m-%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName+".log"),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	log.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	}))

	return log
}

func getInfo() *logrus.Fields {
	pc := make([]uintptr, 5)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	filename, line := f.FileLine(pc[0])
	return &logrus.Fields{
		"file": strings.Replace(filename, repoPath, "", 1),
		"line": line - 1,
		"func": f.Name(),
	}
}

// Info 信息
func Info(args ...interface{}) {
	log.WithFields(*getInfo()).Info(args...)
}

// Fatal
func Fatal(args ...interface{}) {
	log.WithFields(*getInfo()).Fatal(args...)
}

// Debug
func Debug(args ...interface{}) {
	log.WithFields(*getInfo()).Debug(args...)
}

// Warn
func Warn(args ...interface{}) {
	log.WithFields(*getInfo()).Warn(args...)
}

// Error
func Error(args ...interface{}) {
	log.WithFields(*getInfo()).Error(args...)
}

// Panic
func Panic(args ...interface{}) {
	log.WithFields(*getInfo()).Panic(args...)
}
