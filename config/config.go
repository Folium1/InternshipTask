package config

import (
	dto "book/DTO"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Logger *zap.Logger

func Init() {
	db := ConnectDb()
	err := db.AutoMigrate(&dto.BookDTO{})
	if err != nil {
		log.Println(err)
	}

	// Create a new logger configuration
	config := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel), // Set log level to DebugLevel
		OutputPaths:      []string{"stdout", "logs.log"},       // Write logs to stdout and logs.log
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			TimeKey: "time",

			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	// Modify the logger output format
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.Development = true
	config.EncoderConfig.CallerKey = "caller"
	w, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	logger, err := config.Build(zap.ErrorOutput(w))
	if err != nil {
		log.Fatalf("Failed to build logger: %v", err)
	}
	Logger = logger
}

func ConnectDb() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	mysqlConn := os.Getenv("MYSQL")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               mysqlConn,
		DefaultStringSize: 50,
	}), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return db
}
