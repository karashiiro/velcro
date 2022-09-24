package db

import (
	"context"
	"fmt"

	"github.com/velcro-xiv/velcro/ent"
)

func createMessage(ctx context.Context, client *ent.Client, sniff *SniffRecord) (*ent.Message, error) {
	return client.Message.Create().
		SetTimestamp(sniff.Timestamp).
		SetVersion(sniff.Version).
		SetSegment(sniff.Segment).
		SetNillableOpcode(sniff.Opcode).
		SetSourceAddress(sniff.GetSourceAddress().String()).
		SetSourcePort(sniff.SourcePort).
		SetDestinationAddress(sniff.GetDestinationAddress().String()).
		SetDestinationPort(sniff.DestinationPort).
		SetData(sniff.Data).
		Save(ctx)
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

func (a *Archiver) Process() {
	go func() {
	outer:
		for {
			select {
			case <-a.stop:
				break outer
			case sr := <-a.q:
				_, err := createMessage(context.Background(), a.client, sr)
				if err != nil {
					a.logger.LogError(context.Background(), fmt.Sprintf("failed to store record: %v\n", err))
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
