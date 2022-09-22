package main

import (
	"bufio"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"time"

	"entgo.io/ent/dialect"
	"github.com/pkg/errors"
	"github.com/velcro-xiv/velcro/ent"

	"modernc.org/sqlite"
)

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		conn.Close()
		return nil, errors.Wrap(err, "failed to enable enable foreign keys")
	}
	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}

type SniffRecord struct {
	Timestamp          time.Time `json:"t"`
	Version            int       `json:"v"`
	Segment            int       `json:"segment"`
	Opcode             *int      `json:"opcode"`
	SourceAddress      string    `json:"src_addr"`
	SourcePort         int       `json:"src_port"`
	DestinationAddress string    `json:"dst_addr"`
	DestinationPort    int       `json:"dst_port"`
	Data               []byte    `json:"data"`
}

func (s *SniffRecord) GetSourceAddress() net.IP {
	return net.ParseIP(s.SourceAddress)
}

func (s *SniffRecord) GetDestinationAddress() net.IP {
	return net.ParseIP(s.DestinationAddress)
}

func CreateMessage(ctx context.Context, client *ent.Client, sniff *SniffRecord) (*ent.Message, error) {
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

func main() {
	client, err := ent.Open(dialect.SQLite, "file:velcro.db?cache=shared")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v\n", err)
	}
	defer client.Close()

	log.Println("connected to sqlite client")
	log.Println("executing auto-migrations")

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v\n", err)
	}

	// Store data in the database.
	reader := bufio.NewReader(os.Stdin)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatalf("failed reading standard input: %v\n", err)
			}

			continue
		}

		log.Println("found new data")

		sniff := SniffRecord{}
		err = json.Unmarshal(buf, &sniff)
		if err != nil {
			log.Fatalf("failed to unmarshal record: %v\n", err)
		}

		log.Printf("parsed new data: %v\n", &sniff)

		message, err := CreateMessage(context.Background(), client, &sniff)
		if err != nil {
			log.Printf("failed to store record: %v\n", err)
		}

		log.Println("message was created: ", message)
	}
}
