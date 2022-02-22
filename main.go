package main

import (
	"flag"

	db "github.com/eminmuhammadi/memorydb/db"
	tcp "github.com/eminmuhammadi/memorydb/tcp"
)

var ip = flag.String("ip", "0.0.0.0", "--ip 0.0.0.0")
var port = flag.String("port", "8080", "--port 8080")
var timezone = flag.String("tz", "", "--tz \"Asia/Baku\"")

func main() {
	flag.Parse()

	// App structure
	sql := db.Init(*timezone)

	// Migrate database
	if err := db.Migrate(sql); err != nil {
		panic(err)
	}

	// Start TCP server
	tcp.Start(*ip, *port, sql)

	defer db.Close(sql)
}
