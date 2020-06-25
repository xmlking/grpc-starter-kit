package grpc

import (
	"fmt"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/grpclog"
)

type geroLogger struct {
	log *zerolog.Logger
}

func New(log *zerolog.Logger) grpclog.LoggerV2 {
	return &geroLogger{log: log}
}

func (l *geroLogger) Fatal(args ...interface{}) {
	l.log.Fatal().Msg(fmt.Sprint(args...))
}

func (l *geroLogger) Fatalf(format string, args ...interface{}) {
	l.log.Fatal().Msg(fmt.Sprintf(format, args...))
}

func (l *geroLogger) Fatalln(args ...interface{}) {
	l.Fatal(args...)
}

func (l *geroLogger) Error(args ...interface{}) {
	l.log.Error().Msg(fmt.Sprint(args...))
}

func (l *geroLogger) Errorf(format string, args ...interface{}) {
	l.log.Error().Msg(fmt.Sprintf(format, args...))
}

func (l *geroLogger) Errorln(args ...interface{}) {
	l.Error(args...)
}

func (l *geroLogger) Info(args ...interface{}) {
	l.log.Info().Msg(fmt.Sprint(args...))
}

func (l *geroLogger) Infof(format string, args ...interface{}) {
	l.log.Info().Msg(fmt.Sprintf(format, args...))
}

func (l *geroLogger) Infoln(args ...interface{}) {
	l.Info(args...)
}

func (l *geroLogger) Warning(args ...interface{}) {
	l.log.Warn().Msg(fmt.Sprint(args...))
}

func (l *geroLogger) Warningf(format string, args ...interface{}) {
	l.log.Warn().Msg(fmt.Sprintf(format, args...))
}

func (l *geroLogger) Warningln(args ...interface{}) {
	l.Warning(args...)
}

func (l *geroLogger) Print(args ...interface{}) {
	l.log.Info().Msg(fmt.Sprint(args...))
}

func (l *geroLogger) Printf(format string, args ...interface{}) {
	l.log.Info().Msg(fmt.Sprintf(format, args...))
}

func (l *geroLogger) Println(args ...interface{}) {
	l.Print(args...)
}

// grpc logs also follow supplied zerolog logger's level
func (l *geroLogger) V(verbosity int) bool {
	return true
}
