package db

import (
	"context"
	"time"

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

func (l *Logger) LogInfo(ctx context.Context, message string) error {
	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Info, message)
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) LogWarning(ctx context.Context, message string) error {
	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Warning, message)
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) LogError(ctx context.Context, message string) error {
	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Error, message)
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) LogDebug(ctx context.Context, message string) error {
	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Debug, message)
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) LogTrace(ctx context.Context, message string) error {
	_, err := createLogEvent(ctx, l.client, time.Now().UTC(), Trace, message)
	if err != nil {
		return err
	}

	return nil
}
