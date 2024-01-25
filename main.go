package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"os"

	rlog "github.com/sirupsen/logrus"
	sf "github.com/snowflakedb/gosnowflake"

	"github.com/apache/arrow/go/v14/arrow/memory"
)

func main() {
	mylog := sf.GetLogger()
	mylog.SetLogLevel("debug")
	mylog.SetOutput(os.Stdout)
	sf.SetLogger(&mylog)

	ctx := sf.WithArrowAllocator(
		sf.WithArrowBatches(context.Background()),
		memory.DefaultAllocator,
	)

	// insert your dsn here
	dsn := ``
	if len(os.Args) >= 2 {
		dsn = os.Args[1]
	}

	db, err := sql.Open(`snowflake`, dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	conn, err := db.Conn(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var rows driver.Rows
	err = conn.Raw(func(x interface{}) error {
		rows, err = x.(driver.QueryerContext).QueryContext(ctx, `SELECT 1 WHERE 1 = 0`, nil)
		return err
	})
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// actual: panic here
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1025fb9]

	// goroutine 1 [running]:
	// github.com/snowflakedb/sf.(*snowflakeChunkDownloader).getArrowBatches(0xc0003ebe20?)
	// 		github.com/snowflakedb/gosnowflake@v1.7.2/chunk_downloader.go:252 +0x19
	// github.com/snowflakedb/sf.(*snowflakeRows).GetArrowBatches(0xc0000d05b0?)
	// 		github.com/snowflakedb/gosnowflake@v1.7.2/rows.go:171 +0x162
	// main.main()
	// 		github.com/vtan-fortinet/musical-octo-funicular/main.go:59 +0x27a
	batches, err := rows.(sf.SnowflakeRows).GetArrowBatches()
	if err != nil {
		panic(err)
	}

	// expected: len 0
	fmt.Println(`we happy, len`, len(batches))
}

type testLogger struct {
	*rlog.Logger
}

func (log *testLogger) SetLogLevel(level string) error {
	actualLevel, err := rlog.ParseLevel(level)
	if err != nil {
		return err
	}
	log.Level = actualLevel
	return nil
}

func createTestLogger() testLogger {
	var logging = testLogger{rlog.New()}
	var formatter = rlog.JSONFormatter{CallerPrettyfier: sf.SFCallerPrettyfier}
	logging.SetReportCaller(true)
	logging.SetFormatter(&formatter)
	return logging
}
