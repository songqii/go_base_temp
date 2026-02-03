package templates

import "fmt"

func MainGo(moduleName string) string {
	return fmt.Sprintf(`package main

import (
	"flag"

	"%s/conf"
	"%s/log"
	"%s/pkg/cache"
	"%s/pkg/controller"
	"%s/pkg/storage"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config_path", "dev.yaml", "config file path")
	flag.Parse()
}

func main() {
	// Load config
	if err := conf.InitConfig(configPath); err != nil {
		panic(err)
	}

	// Initialize logger
	log.ZapInit()
	log.Zlog.Info("server init")
	log.Zlog.Infof("config path: %%s", configPath)
	log.Zlog.Infof("env: %%s", conf.Conf.Env)

	// Initialize cache
	if err := cache.NewCache(); err != nil {
		log.Zlog.Fatalf("cache init err: %%+v", err)
	}

	// Initialize database
	storage.Init()

	// Start API server
	controller.StartApi()
}
`, moduleName, moduleName, moduleName, moduleName, moduleName)
}

func ConfGo(moduleName string) string {
	return `package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Conf       *Config
	configType = "yaml"
)

type Config struct {
	Debug     bool
	Mode      string
	Env       string
	MysqlAddr string

	Socket struct {
		Port string
	}

	Redis struct {
		Addr string
		Pwd  string
		Db   int
	}

	Log struct {
		LogPath  string
		LogName  string
		LogLevel string
	}

	JWT struct {
		Secret string
		Expire int
	}
}

func InitConfig(path string) error {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType(configType)

	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(&Conf); err != nil {
		return err
	}
	fmt.Printf("Config loaded: %+v\n", Conf)
	return nil
}
`
}

func LogGo(moduleName string) string {
	return fmt.Sprintf(`package log

import (
	"%s/conf"
	"io"
	"os"
	"strings"
	"time"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zlog *zap.SugaredLogger

var logMap = map[string]zapcore.Level{
	"Debug": zapcore.DebugLevel,
	"Info":  zapcore.InfoLevel,
	"Warn":  zapcore.WarnLevel,
	"Error": zapcore.ErrorLevel,
}

func SetTimer(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func ZapInit() {
	if conf.Conf.Debug {
		debugInit()
		return
	}

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
		TimeKey:      "ts",
		EncodeTime:   SetTimer,
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if val, ok := logMap[conf.Conf.Log.LogLevel]; ok {
			return lvl >= val
		}
		return lvl >= zapcore.DebugLevel
	})

	writer := zapcore.AddSync(getWriter(conf.Conf.Log.LogPath + "log_" + conf.Conf.Log.LogName))
	core := zapcore.NewTee(zapcore.NewCore(encoder, writer, lowPriority))
	log := zap.New(core, zap.AddCaller())
	Zlog = log.Sugar()
}

func getWriter(filename string) io.Writer {
	hook, err := rotateLogs.New(
		strings.Replace(filename, ".log", "", -1)+"_%%Y_%%m_%%d_%%H.log",
		rotateLogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func debugInit() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     SetTimer,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level),
	)
	log := zap.New(core, zap.AddCaller())
	Zlog = log.Sugar()
}
`, moduleName)
}
