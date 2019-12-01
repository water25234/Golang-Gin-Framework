package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var hook lumberjack.Logger
var encoderConfig zapcore.EncoderConfig

func init() {
	hook = lumberjack.Logger{
		Filename:   "./logs/Golang-Gin-Framework.log", // 日誌檔案路徑
		MaxSize:    128,                               // 每個日誌檔案儲存的最大尺寸 單位：M
		MaxBackups: 30,                                // 日誌檔案最多儲存多少個備份
		MaxAge:     7,                                 // 檔案最多儲存多少天
		Compress:   true,                              // 是否壓縮
	}

	encoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小寫編碼器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 時間格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路徑編碼器
		EncodeName:     zapcore.FullNameEncoder,
	}
}

func Info(logContext string) {
	// 設定日誌級別
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 編碼器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 列印到控制檯和檔案
		atomicLevel, // 日誌級別
	)

	// 開啟開發模式，堆疊跟蹤
	caller := zap.AddCaller()
	// 開啟檔案及行號
	development := zap.Development()
	// 設定初始化欄位
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 構造日誌
	logger := zap.New(core, caller, development, filed)

	logger.Info(logContext)

	// logger.Info("無法獲取網址",
	//	zap.String("url", "http://www.baidu.com"),
	//	zap.Int("attempt", 3),
	//	zap.Duration("backoff", time.Second))
}
