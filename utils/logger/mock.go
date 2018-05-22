package logger

import (
	"github.com/maraino/go-mock"
)

type MockLogger struct {
	mock.Mock
}

func (l *MockLogger) Fatal(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Fatalf(format string, args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Fatalln(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Panic(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Panicf(format string, args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Panicln(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Print(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Printf(format string, args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Println(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Debug(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Debugf(format string, args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Debugln(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Error(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Errorf(format string, args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Errorln(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Info(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Infof(format string, args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Infoln(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Warn(args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Warnf(format string, args ...interface{}) {
	l.Called()
}

func (l *MockLogger) Warnln(args ...interface{}) {
	l.Called()
}