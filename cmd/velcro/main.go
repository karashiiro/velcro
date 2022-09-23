package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"entgo.io/ent/dialect"
	"github.com/velcro-xiv/velcro/db"
	_ "github.com/velcro-xiv/velcro/driver"
	"github.com/velcro-xiv/velcro/ent"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:velcro.db?cache=shared&journal_mode=WAL")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed opening connection to sqlite: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	logger := db.NewLogger(client)

	logger.LogInfo(context.Background(), "connected to sqlite client")
	logger.LogInfo(context.Background(), "executing auto-migrations")

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.LogError(context.Background(), fmt.Sprintf("failed creating schema resources: %v\n", err))
	}

	// Store data in the database.
	archiver := db.NewArchiver(client, logger)
	archiver.Process()
	defer archiver.Stop()

	reader := bufio.NewReader(os.Stdin)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				logger.LogError(context.Background(), fmt.Sprintf("failed reading standard input: %v\n", err))
			}

			continue
		}

		logger.LogDebug(context.Background(), "found new data")

		sniff := db.SniffRecord{}
		err = json.Unmarshal(buf, &sniff)
		if err != nil {
			logger.LogError(context.Background(), fmt.Sprintf("failed to unmarshal record: %v\n", err))
		}

		logger.LogDebug(context.Background(), fmt.Sprintf("parsed new data: %v\n", &sniff))

		archiver.Store(&sniff)
	}
}
