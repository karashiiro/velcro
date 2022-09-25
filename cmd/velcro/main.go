package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"entgo.io/ent/dialect"
	"github.com/pkg/errors"
	"github.com/velcro-xiv/velcro/db"
	_ "github.com/velcro-xiv/velcro/driver"
	"github.com/velcro-xiv/velcro/ent"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:velcro.db?_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", errors.Wrap(err, "failed opening connection to sqlite"))
		os.Exit(1)
	}
	defer client.Close()

	logger := db.NewLogger(client)

	logger.LogInfo(context.Background(), "connected to sqlite client")
	logger.LogInfo(context.Background(), "executing auto-migrations")

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.LogError(context.Background(), fmt.Sprintf("%v", errors.Wrap(err, "failed creating schema resources")))
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
				logger.LogError(context.Background(), fmt.Sprintf("%v", errors.Wrap(err, "failed reading standard input")))
			}

			continue
		}

		_, err = fmt.Print(string(buf))
		if err != nil {
			logger.LogError(context.Background(), fmt.Sprintf("%v", errors.Wrap(err, "failed to print data line to standard output")))
			continue
		}

		logger.LogDebug(context.Background(), "found new data")

		sniff := db.SniffRecord{}
		err = json.Unmarshal(buf, &sniff)
		if err != nil {
			logger.LogError(context.Background(), fmt.Sprintf("%v", errors.Wrap(err, "failed to unmarshal record")))
			continue
		}

		if sniff.Version != 2 {
			logger.LogError(context.Background(), fmt.Sprintf("record version is unsupported: %v", &sniff))
			continue
		}

		logger.LogDebug(context.Background(), fmt.Sprintf("parsed new data: %v", &sniff))

		archiver.Store(&sniff)
	}
}
