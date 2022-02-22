package main

import (
	config "github.com/eminmuhammadi/memorydb/config"
	db "github.com/eminmuhammadi/memorydb/db"
	tcp "github.com/eminmuhammadi/memorydb/tcp"
)

var conf = config.GetConfig()

func main() {
	// App structure
	sql := db.Init(*&conf.TimeZone)

	// Migrate database
	if err := db.Migrate(sql); err != nil {
		panic(err)
	}

	// Start TCP server
	tcp.Start(*&conf.IP, *&conf.PORT, sql)

	defer db.Close(sql)
}
