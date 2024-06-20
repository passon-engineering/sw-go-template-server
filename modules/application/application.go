package application

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/tpasson/sw-go-logger-lib/logger"
	"github.com/tpasson/sw-go-utility-lib/networking"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Application struct {
	ServerPath  string
	SystemIP    string
	Logger      *logger.Logger
	Config      Config
	MongoClient *mongo.Client
}

type Config struct {
	HttpAddress  string
	HttpsAddress string
	WebTlsCert   string
	WebTlsKey    string
	WebDirectory string
	LogDirectory string
	MongoURL     string
}

func Init(config Config) *Application {
	startTime := time.Now()
	var err error

	app := &Application{
		ServerPath: filepath.Dir(os.Args[0]),
		SystemIP:   "",
		Logger:     &logger.Logger{},
		Config:     config,
	}

	// Determine if logs should be written to a file based on the LogDirectory configuration
	outputToFile := app.Config.LogDirectory != ""
	logOutputDefinition := "Logging to files disabled - none or invalid path provided."

	if outputToFile {
		logOutputDefinition = "Writing log files to: " + app.Config.LogDirectory
	}

	app.Logger, err = logger.NewLogger(
		[]logger.LogFormat{
			logger.FORMAT_TIMESTAMP,
			logger.FORMAT_STATUS,
			logger.FORMAT_INFO,
			logger.FORMAT_PRE_TEXT,
			logger.FORMAT_HTTP_REQUEST,
			logger.FORMAT_ID,
			logger.FORMAT_SOURCE,
			logger.FORMAT_DATA,
			logger.FORMAT_ERROR,
			logger.FORMAT_PROCESSING_TIME,
		}, logger.Options{
			OutputToStdout:   true,
			OutputToFile:     outputToFile,
			OutputFolderPath: app.Config.LogDirectory,
		}, logger.Container{
			Status: logger.STATUS_INFO,
			Info:   "System Logger successfully started! Awaiting logger tasks. " + logOutputDefinition,
		})
	if err != nil {
		log.Fatalf("Could not initialize logger: %v", err)
	}

	app.Logger.Entry(logger.Container{
		Status:         logger.STATUS_INFO,
		Info:           "Server path: " + app.ServerPath,
		ProcessingTime: time.Since(startTime),
	})

	ip, err := networking.GetNetworkExternalIP()
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status: logger.STATUS_ERROR,
			Error:  "Could not get network external IP: " + err.Error(),
		})
		os.Exit(1)
	}
	app.Logger.Entry(logger.Container{
		Status: logger.STATUS_INFO,
		Info:   "Network external IP: " + ip,
	})

	app.SystemIP = ip

	// Initialize MongoDB client
	clientOptions := options.Client().ApplyURI(app.Config.MongoURL)
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status: logger.STATUS_ERROR,
			Error:  "Could not connect to MongoDB: " + err.Error(),
		})
	} else {
		// Test the MongoDB connection
		err = mongoClient.Ping(context.TODO(), readpref.Primary())
		if err != nil {
			app.Logger.Entry(logger.Container{
				Status: logger.STATUS_ERROR,
				Error:  "Could not ping MongoDB: " + err.Error(),
			})
		} else {
			app.MongoClient = mongoClient
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           "Successfully connected to MongoDB",
				ProcessingTime: time.Since(startTime),
			})
		}
	}

	app.Logger.Entry(logger.Container{
		Status:         logger.STATUS_INFO,
		Info:           "Basic app framework successfully initialized",
		ProcessingTime: time.Since(startTime),
	})

	return app
}
