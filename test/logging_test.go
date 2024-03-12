package test

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

//import go get github.com/sirupsen/logrus

func TestLogger(t *testing.T) {
	logger := logrus.New()
	logger.Println("hello Logger")
	fmt.Println("hello logger")
}

func TestLoggerLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.Trace("This is Trace")  //Trace
	logger.Debugf("This is Debug") //Debug
	logger.Info("This is Info")    //Info
	logger.Warn("This is Warn")    //Warn
	logger.Error("This is Error")  //Error
	//Fatal
	//Panic
}

func TestLoggerOutput(t *testing.T) {
	logger := logrus.New()
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)
	logger.Info("hello logging")
	logger.Warn("hello warning logging")
	logger.Error("hello error logging")
}

func TestLoggerOutputjsonFormat(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	logger.Info("hello logging")
	logger.Warn("hello warning logging")
	logger.Error("hello error logging")
}

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("username", "iqbal").Info("hello world")

	logger.WithField("username", "iqbal").WithField("name", "Iqbal Fauzan").Info("Info masse")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "iqbalfauzan",
		"name":     "Iqbal Fauzan Name",
	}).Infof("hello world")
}

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample hook", entry.Level, entry.Message) // execute send email / send notif about log
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Infof("hello info")

	logger.Warn("hello warning")
}
