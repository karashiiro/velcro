package main

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	"github.com/karashiiro/velcro/ent"

	_ "github.com/mattn/go-sqlite3"
)

type SniffRecord struct {
	Timestamp          *time.Time `json:"t"`
	Version            int        `json:"v"`
	SourceAddress      string     `json:"src_addr"`
	SourcePort         int        `json:"src_port"`
	DestinationAddress string     `json:"dst_addr"`
	DestinationPort    int        `json:"dst_port"`
	Data               []byte     `json:"data"`
}

func (s *SniffRecord) GetSourceAddress() net.IP {
	return net.ParseIP(s.SourceAddress)
}

func (s *SniffRecord) GetDestinationAddress() net.IP {
	return net.ParseIP(s.DestinationAddress)
}

func CreateMessage(ctx context.Context, client *ent.Client, sniff *SniffRecord) (*ent.Message, error) {
	return client.Message.Create().
		SetTimestamp(*sniff.Timestamp).
		SetVersion(sniff.Version).
		SetSourceAddress(sniff.GetSourceAddress().String()).
		SetSourcePort(sniff.SourcePort).
		SetDestinationAddress(sniff.GetDestinationAddress().String()).
		SetDestinationPort(sniff.DestinationPort).
		SetData(sniff.Data).
		Save(ctx)
}

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	log.Println("connected to sqlite client")
	log.Println("executing auto-migrations")

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Store data in the database.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		log.Println("reading data")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				log.Fatalf("failed reading standard input: %v", err)
			}

			continue
		}

		log.Println("found new data")

		sniff := SniffRecord{}
		err := json.Unmarshal(scanner.Bytes(), &sniff)
		if err != nil {
			log.Fatalf("failed to unmarshal record: %v", err)
		}

		log.Printf("parsed new data: %v\n", &sniff)

		message, err := CreateMessage(context.Background(), client, &sniff)
		if err != nil {
			log.Fatalf("failed to store record: %v", err)
		}

		log.Println("message was created: ", message)
	}
}
