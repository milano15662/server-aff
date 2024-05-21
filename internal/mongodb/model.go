package mongodb

import "go.mongodb.org/mongo-driver/event"

// ConnectTLSOpts ...
type ConnectTLSOpts struct {
	ReplSet             string
	CaFile              string
	CertKeyFile         string
	CertKeyFilePassword string
	ReadPreferenceMode  string
}

// ConnectStandaloneOpts ...
type ConnectStandaloneOpts struct {
	AuthMechanism string
	AuthSource    string
	Username      string
	Password      string
}

// Config ...
type Config struct {
	Host    string
	DBName  string
	Monitor *event.CommandMonitor

	TLS        *ConnectTLSOpts
	Standalone *ConnectStandaloneOpts
}
