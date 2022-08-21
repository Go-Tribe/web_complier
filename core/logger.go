package core

import (
	"io"
	"path/filepath"
	"time"
	c "web_complier/configs"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZLogger *zap.Logger

func InitLogger() {
	ZLogger = createZapLog()
}

// initZapLog 初始化 zap 日志
func createZapLog() *zap.Logger {
	// 开启 debug
	if c.Config.Debug == true {
		if Logger, err := zap.NewDevelopment(); err == nil {
			return Logger
		} else {
			panic("创建zap日志包失败，详情：" + err.Error())
		}
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	filename := filepath.Join(c.Config.StaticBasePath, "/logs/", c.Config.Logger.Filename)
	var writer zapcore.WriteSyncer
	if c.Config.Logger.DefaultDivision == "size" {
		// 按文件大小切割日志
		writer = zapcore.AddSync(getLumberJackWriter(filename))
	} else {
		// 按天切割日志
		writer = zapcore.AddSync(getRotateWriter(filename))
	}

	zapCore := zapcore.NewCore(encoder, writer, zap.InfoLevel)
	// zap.AddStacktrace(zap.WarnLevel)
	return zap.New(zapCore, zap.AddCaller())
}

// getRotateWriter 按日期切割日志
func getRotateWriter(filename string) io.Writer {
	maxAge := time.Duration(c.Config.Logger.DivisionTime.MaxAge)
	rotationTime := time.Duration(c.Config.Logger.DivisionTime.RotationTime)
	hook, err := rotatelogs.New(
		filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*maxAge),
		rotatelogs.WithRotationTime(time.Hour*rotationTime), // 默认一天
	)

	if err != nil {
		panic(err)
	}

	return hook
}

// getLumberJackWriter 按文件切割日志
func getLumberJackWriter(filename string) io.Writer {
	// 日志切割配置
	return &lumberjack.Logger{
		Filename:   filename,                                // 日志文件的位置
		MaxSize:    c.Config.Logger.DivisionSize.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: c.Config.Logger.DivisionSize.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     c.Config.Logger.DivisionSize.MaxAge,     // 保留旧文件的最大天数
		Compress:   c.Config.Logger.DivisionSize.Compress,   // 是否压缩/归档旧文件
	}
}
