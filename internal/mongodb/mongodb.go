package mongodb

import (
	"context"
	"fmt"
	"git.selly.red/Cashbag-B2B/server-aff/internal/config"
	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
)

var db *mongo.Database

// Connect ...
func Connect(cfg config.MongoDBCfg) {
	var (
		err error
		tls *ConnectTLSOpts
	)

	if cfg.ReplicaSet != "" {
		tls = &ConnectTLSOpts{
			ReplSet:             cfg.ReplicaSet,
			CaFile:              cfg.CAPem,
			CertKeyFile:         cfg.CertPem,
			CertKeyFilePassword: cfg.CertKeyFilePassword,
			ReadPreferenceMode:  cfg.ReadPrefMode,
		}
	}

	// Connect
	db, err = connect(Config{
		Host:       cfg.URI,
		DBName:     cfg.DBName,
		Monitor:    apmmongo.CommandMonitor(),
		TLS:        tls,
		Standalone: &ConnectStandaloneOpts{},
	})
	if err != nil {
		panic(err)
	}

	// index
	go colIndexes()
}

//
// method private
//

func connect(cfg Config) (*mongo.Database, error) {
	if cfg.TLS != nil && cfg.TLS.ReplSet != "" {
		return connectWithTLS(cfg)
	}
	connectOptions := options.ClientOptions{}
	opts := cfg.Standalone

	// Set auth if existed
	if opts != nil && opts.Username != "" && opts.Password != "" {
		connectOptions.Auth = &options.Credential{
			AuthMechanism: opts.AuthMechanism,
			AuthSource:    opts.AuthSource,
			Username:      opts.Username,
			Password:      opts.Password,
		}
	}
	if cfg.Monitor != nil {
		connectOptions.SetMonitor(cfg.Monitor)
	}

	// Connect
	client, err := mongo.Connect(context.Background(), connectOptions.ApplyURI(cfg.Host))
	if err != nil {
		return nil, err
	}
	fmt.Printf("⚡️[mongodb]: connected to %s/%s \n", cfg.Host, cfg.DBName)

	// send a ping to confirm a successful connection
	if err = client.Database("admin").RunCommand(context.Background(),
		bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	// Set data
	db = client.Database(cfg.DBName)
	return db, nil
}

func connectWithTLS(cfg Config) (*mongo.Database, error) {
	ctx := context.Background()
	opts := cfg.TLS

	caFile, err := initFileFromBase64String("ca.pem", opts.CaFile)
	if err != nil {
		return nil, err
	}
	certFile, err := initFileFromBase64String("cert.pem", opts.CertKeyFile)
	if err != nil {
		return nil, err
	}
	pwd := base64DecodeToString(opts.CertKeyFilePassword)
	uri := getURIWithTLS(cfg, caFile.Name(), certFile.Name(), pwd)
	readPref := getReadPref(opts.ReadPreferenceMode)
	clientOpts := options.Client().SetReadPreference(readPref).SetReplicaSet(opts.ReplSet).ApplyURI(uri)
	if cfg.Monitor != nil {
		clientOpts.SetMonitor(cfg.Monitor)
	}
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.SecondaryPreferred()); err != nil {
		return nil, err
	}
	db = client.Database(cfg.DBName)

	fmt.Printf("⚡️[mongodb/tls]: connected to %s/%s \n", cfg.Host, cfg.DBName)
	return db, err
}

func getURIWithTLS(cfg Config, caFilePath, certFilePath, pwd string) string {
	host := cfg.Host
	if strings.Contains(host, "?") {
		host += "&"
	} else {
		if !strings.HasSuffix(host, "/") {
			host += "/?"
		} else {
			host += "?"
		}
	}
	s := "%stls=true&tlsCAFile=./%s&tlsCertificateKeyFile=./%s&tlsCertificateKeyFilePassword=%s&authMechanism=MONGODB-X509"
	uri := fmt.Sprintf(s, host, caFilePath, certFilePath, pwd)
	return uri
}

func getReadPref(mode string) *readpref.ReadPref {
	m, err := readpref.ModeFromString(mode)
	if err != nil {
		m = readpref.SecondaryPreferredMode
	}
	readPref, err := readpref.New(m)
	if err != nil {
		fmt.Println("mongodb.getReadPref err: ", err, m)
	}
	return readPref
}
