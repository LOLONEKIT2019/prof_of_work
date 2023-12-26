package config

import "time"

var Config = struct {
	ZerosCount int
	MaxNonce   int
	ServerHost string
	ServerPort int
	TaskTTL    time.Duration
}{
	ZerosCount: 4,
	MaxNonce:   1000000000,
	ServerHost: "server",
	ServerPort: 8080,
	TaskTTL:    time.Minute,
}
