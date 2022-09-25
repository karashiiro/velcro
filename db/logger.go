package db

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/velcro-xiv/velcro/ent"
)

type LogLevel int

const (
	Info LogLevel = iota
	Warning
	Error
	Debug
	Trace
)

func createLogEvent(ctx context.Context, client *ent.Client, t time.Time, level LogLevel, message string) (*ent.LogEvent, error) {
	return client.LogEvent.Create().
		SetTimestamp(t).
		SetLevel(int(level)).
		SetMessage(message).
		Save(ctx)
}

type Logger struct {
	client *ent.Client
}

func NewLogger(client *ent.Client) *Logger {
	return &Logger{
		client: client,
	}
}

func (l *Logger) LogInfo(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Info, message)
	if err != nil {
		return errors.Wrap(err, "failed to log info message")
	}

	return nil
}

func (l *Logger) LogWarning(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Warning, message)
	if err != nil {
		return errors.Wrap(err, "failed to log warning message")
	}

	return nil
}

func (l *Logger) LogError(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Error, message)
	if err != nil {
		return errors.Wrap(err, "failed to log error message")
	}

	return nil
}

func (l *Logger) LogDebug(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Debug, message)
	if err != nil {
		return errors.Wrap(err, "failed to log debug message")
	}

	return nil
}

func (l *Logger) LogTrace(message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Trace, message)
	if err != nil {
		return errors.Wrap(err, "failed to log trace message")
	}

	return nil
}
