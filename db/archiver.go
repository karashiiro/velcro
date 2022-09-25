package db

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/velcro-xiv/velcro/ent"
)

func createMessage(ctx context.Context, client *ent.Client, sniff *SniffRecord) (*ent.Message, error) {
	create := client.Message.Create().
		SetTimestamp(sniff.Timestamp).
		SetVersion(sniff.Version).
		SetSourceAddress(sniff.SourceAddress).
		SetSourcePort(sniff.SourcePort).
		SetDestinationAddress(sniff.DestinationAddress).
		SetDestinationPort(sniff.DestinationPort).
		SetSize(sniff.SegmentHeader.Size).
		SetSourceActor(sniff.SegmentHeader.SourceActor).
		SetTargetActor(sniff.SegmentHeader.TargetActor).
		SetSegmentType(sniff.SegmentHeader.Type).
		SetData(sniff.MessageData)
	if sniff.MessageHeader != nil {
		create.
			SetNillableOpcode(&sniff.MessageHeader.Opcode).
			SetNillableServer(&sniff.MessageHeader.Server).
			SetNillableTimestampRaw(&sniff.MessageHeader.Timestamp)
	}
	return create.Save(ctx)
}

type Archiver struct {
	client  *ent.Client
	logger  *Logger
	q       chan *SniffRecord
	stop    chan bool
	stopped chan bool
}

func NewArchiver(client *ent.Client, logger *Logger) *Archiver {
	return &Archiver{
		client:  client,
		logger:  logger,
		q:       make(chan *SniffRecord),
		stop:    make(chan bool, 1),
		stopped: make(chan bool, 1),
	}
}

func (a *Archiver) Store(sr *SniffRecord) {
	a.q <- sr
}

func (a *Archiver) storeRecord(ctx context.Context, sr *SniffRecord) error {
	_, err := createMessage(ctx, a.client, sr)
	if err != nil {
		return errors.Wrap(err, "failed to store record")
	}

	return nil
}

func (a *Archiver) Process() {
	go func() {
	outer:
		for {
			select {
			case <-a.stop:
				break outer
			case sr := <-a.q:
				err := a.storeRecord(context.Background(), sr)
				if err != nil {
					a.logger.LogError(fmt.Sprintf("%v", err))
				}
			}
		}

		a.stopped <- true
	}()
}

func (a *Archiver) Stop() {
	a.stop <- true
	<-a.stopped
}
