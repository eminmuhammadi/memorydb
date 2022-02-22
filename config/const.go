package config

import "flag"

type Config struct {
	IP       string
	PORT     string
	TimeZone string
}

var IP = flag.String("ip", "0.0.0.0", "--ip 0.0.0.0")
var PORT = flag.String("port", "8080", "--port 8080")
var TimeZone = flag.String("tz", "", "--tz \"Asia/Baku\"")

func GetConfig() Config {
	flag.Parse()
	return Config{
		IP:       *IP,
		PORT:     *PORT,
		TimeZone: *TimeZone,
	}
}
