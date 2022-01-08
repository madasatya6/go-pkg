package uberzap

import (
	"os"
	"time"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
// logging with a file log
// reference uberzap: https://medium.com/@gustavo.nabakseixas/go-using-uber-zap-in-your-application-135756f23bdc
func createDirectoryIfNotExists() {
	// create log directory
	path, _ := os.Getwd()
	if _, err := os.Stat(fmt.Sprintf("%s/logs", path)); os.IsNotExist(err) {
		_ = os.Mkdir("logs", os.ModePerm)
	}
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(path + "/logs/" + filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
	})
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitLogger(filename string) {
	createDirectoryIfNotExists()
	writerSync := getLogWriter(filename)
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writerSync, zapcore.DebugLevel)
	logg := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(logg)
}

/*
func main(){
	fmt.Println("Implementasi uber zap logging")
	InitLogger("januari-2022.txt")
	zap.S().Info("Information debug uber zap")
}*/