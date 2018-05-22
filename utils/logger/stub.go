package logger

type StubLogger struct {
}

func (l *StubLogger) Fatal(args ...interface{}) {
}

func (l *StubLogger) Fatalf(format string, args ...interface{}) {
}

func (l *StubLogger) Fatalln(args ...interface{}) {
}

func (l *StubLogger) Panic(args ...interface{}) {
}

func (l *StubLogger) Panicf(format string, args ...interface{}) {
}

func (l *StubLogger) Panicln(args ...interface{}) {
}

func (l *StubLogger) Print(args ...interface{}) {
}

func (l *StubLogger) Printf(format string, args ...interface{}) {
}

func (l *StubLogger) Println(args ...interface{}) {
}

func (l *StubLogger) Debug(args ...interface{}) {
}

func (l *StubLogger) Debugf(format string, args ...interface{}) {
}

func (l *StubLogger) Debugln(args ...interface{}) {
}

func (l *StubLogger) Error(args ...interface{}) {
}

func (l *StubLogger) Errorf(format string, args ...interface{}) {
}

func (l *StubLogger) Errorln(args ...interface{}) {
}

func (l *StubLogger) Info(args ...interface{}) {
}

func (l *StubLogger) Infof(format string, args ...interface{}) {
}

func (l *StubLogger) Infoln(args ...interface{}) {
}

func (l *StubLogger) Warn(args ...interface{}) {
}

func (l *StubLogger) Warnf(format string, args ...interface{}) {
}

func (l *StubLogger) Warnln(args ...interface{}) {
}