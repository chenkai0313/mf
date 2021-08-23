package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"api/config"
)

var zapLog *zap.Logger

type Log struct{}

var ZapLog Log

func (log Log) dbLog(lineNum, time, rows, sql interface{}) {
	//是否需要重新初始化，生成新的日志文件
	if !needInitLog() {
		InitLogger()
	}
	msg := fmt.Sprintf("[%vms] [rows:%v] %v", time, rows, sql)
	zapLog.Info(msg,
		zap.String("category", "dbExecLog"),
		zap.String("line", lineNum.(string)),
	)
}

func (log Log) Info(category, msg string) {
	//是否需要重新初始化，生成新的日志文件
	if !needInitLog() {
		InitLogger()
	}
	_, file, line, _ := runtime.Caller(1)
	lineNum := file + ":" + strconv.Itoa(line)
	zapLog.Info(msg,
		zap.String("category", category),
		zap.String("line", lineNum),
	)
}

func (log Log) Warn(category, msg string) {
	//是否需要重新初始化，生成新的日志文件
	if !needInitLog() {
		InitLogger()
	}
	_, file, line, _ := runtime.Caller(1)
	lineNum := file + ":" + strconv.Itoa(line)
	zapLog.Warn(msg,
		zap.String("category", category),
		zap.String("line", lineNum),
	)
}

func (log Log) Error(category, msg string) {
	//是否需要重新初始化，生成新的日志文件
	if !needInitLog() {
		InitLogger()
	}
	_, file, line, _ := runtime.Caller(1)
	lineNum := file + ":" + strconv.Itoa(line)

	zapLog.Error(msg,
		zap.String("category", category),
		zap.String("line", lineNum),
	)
}

func (log Log) Fatal(category, msg string) {
	//是否需要重新初始化，生成新的日志文件
	if !needInitLog() {
		InitLogger()
	}
	_, file, line, _ := runtime.Caller(1)
	lineNum := file + ":" + strconv.Itoa(line)
	zapLog.Fatal(msg,
		zap.String("category", category),
		zap.String("line", lineNum),
	)
}

func (log Log) Debug(category, msg string) {
	//是否需要重新初始化，生成新的日志文件
	if !needInitLog() {
		InitLogger()
	}
	_, file, line, _ := runtime.Caller(1)
	lineNum := file + ":" + strconv.Itoa(line)
	zapLog.Debug(msg,
		zap.String("category", category),
		zap.String("line", lineNum),
	)
}

func (log Log) Panic(category, msg string) {
	//是否需要重新初始化，生成新的日志文件
	if !needInitLog() {
		InitLogger()
	}
	_, file, line, _ := runtime.Caller(1)
	lineNum := file + ":" + strconv.Itoa(line)
	zapLog.Panic(msg,
		zap.String("category", category),
		zap.String("line", lineNum),
	)
}

//设置统一日志路径<elk 会获取此路径日志>
func InitLogger() {
	var logpath = config.GetLogPath()
	timeStr := time.Now().Format("2006-01-02")
	fileName := logpath + "/" + timeStr + ".log"
	existBool, _ := isFileExist(fileName)
	if !existBool {
		//创建目录
		err := os.MkdirAll(logpath, os.ModePerm)
		if err != nil {
			log.Panic("日志文件夹创建失败")
		}
		//创建文件
		f, err := os.Create(fileName)
		if err != nil {
			log.Panic("日志文件创建失败")
		}
		f.Close()
	}

	hook := lumberjack.Logger{
		Filename:   fileName, // 日志文件路径
		MaxSize:    128,      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,       // 日志文件最多保存多少个备份
		MaxAge:     7,        // 文件最多保存多少天
		Compress:   true,     // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,                      // 小写编码器
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,                     //
		EncodeCaller:   zapcore.FullCallerEncoder,                          // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                          // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(Writer{}), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel,                                                                    // 日志级别
	)

	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", config.GetHttpServer().Name))
	// 构造日志
	logger := zap.New(core, filed)

	zapLog = logger
}

//是否需要重新初始化,创建新的日志文件
func needInitLog() bool {
	var logpath = config.GetLogPath()
	timeStr := time.Now().Format("2006-01-02")
	fileName := logpath + "/" + timeStr + ".log"
	existBool, _ := isFileExist(fileName)
	return existBool
}

//判断文件文件夹是否存在
func isFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

func (w Writer) Write(p []byte) (n int, err error) {
	type logStruct struct {
		Level       string `json:"level"`
		Time        string `json:"time"`
		Msg         string `json:"msg"`
		ServiceName string `json:"service_name"`
		Category    string `json:"category"`
		Line        string `json:"line"`
	}
	var tmpLog logStruct
	err = json.Unmarshal(p, &tmpLog)
	if err != nil {
		return 0, err
	}
	str := fmt.Sprintf(`{"level":"%v","time":"%v","msg":"%v","serviceName":"%v","category":"%v","line":" %s "} `, tmpLog.Level, tmpLog.Time, tmpLog.Msg, tmpLog.ServiceName, tmpLog.Category, tmpLog.Line)
	fmt.Println(str)
	return
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	if config.GetDebug() {
		fmt.Printf(format+"\n", args...)
	} else {
		var lineNum, sqlTime, rows, sql interface{}
		for k, v := range args {
			switch k {
			case 0:
				lineNum = v
			case 1:
				sqlTime = v
			case 2:
				rows = v
			case 3:
				sql = v
			}
		}
		ZapLog.dbLog(lineNum, sqlTime, rows, sql)
	}
}
