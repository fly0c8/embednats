package main

import (
	"fmt"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {

	s, err := server.NewServer(&server.Options{
		ConfigFile:                 "",
		ServerName:                 "arnis_server",
		Host:                       "127.0.0.1",
		Port:                       6666,
		ClientAdvertise:            "",
		Trace:                      false,
		Debug:                      false,
		TraceVerbose:               false,
		NoLog:                      false,
		NoSigs:                     false,
		NoSublistCache:             false,
		NoHeaderSupport:            false,
		DisableShortFirstPing:      false,
		Logtime:                    false,
		MaxConn:                    0,
		MaxSubs:                    0,
		Nkeys:                      nil,
		Users:                      nil,
		Accounts:                   nil,
		NoAuthUser:                 "",
		SystemAccount:              "",
		NoSystemAccount:            false,
		AllowNewAccounts:           false,
		Username:                   "",
		Password:                   "",
		Authorization:              "",
		PingInterval:               0,
		MaxPingsOut:                0,
		HTTPHost:                   "",
		HTTPPort:                   0,
		HTTPBasePath:               "",
		HTTPSPort:                  0,
		AuthTimeout:                0,
		MaxControlLine:             0,
		MaxPayload:                 0,
		MaxPending:                 0,
		Cluster:                    server.ClusterOpts{},
		Gateway:                    server.GatewayOpts{},
		LeafNode:                   server.LeafNodeOpts{},
		JetStream:                  false,
		JetStreamMaxMemory:         0,
		JetStreamMaxStore:          0,
		JetStreamDomain:            "",
		JetStreamKey:               "",
		StoreDir:                   "",
		Websocket:                  server.WebsocketOpts{},
		MQTT:                       server.MQTTOpts{},
		ProfPort:                   0,
		PidFile:                    "",
		PortsFileDir:               "",
		LogFile:                    "",
		LogSizeLimit:               0,
		Syslog:                     false,
		RemoteSyslog:               "",
		Routes:                     nil,
		RoutesStr:                  "",
		TLSTimeout:                 0,
		TLS:                        false,
		TLSVerify:                  false,
		TLSMap:                     false,
		TLSCert:                    "",
		TLSKey:                     "",
		TLSCaCert:                  "",
		TLSConfig:                  nil,
		TLSPinnedCerts:             nil,
		AllowNonTLS:                false,
		WriteDeadline:              0,
		MaxClosedClients:           0,
		LameDuckDuration:           0,
		LameDuckGracePeriod:        0,
		MaxTracedMsgLen:            0,
		TrustedKeys:                nil,
		TrustedOperators:           nil,
		AccountResolver:            nil,
		AccountResolverTLSConfig:   nil,
		CustomClientAuthentication: nil,
		CustomRouterAuthentication: nil,
		CheckConfig:                false,
		ConnectErrorReports:        0,
		ReconnectErrorReports:      0,
		Tags:                       nil,
		OCSPConfig:                 nil,
	})
	if err != nil {
		log.Fatal(err)
	}
	s.Start()
	log.Println("server is running:", s.Running())

	opts := []nats.Option{nats.Name("NATS Sample Publisher")}
	nc, err := nats.Connect("nats://127.0.0.1:6666", opts...)
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	nc.Subscribe("from_msgbus", fromMsgbus)
	nc.Subscribe("to_msgbus", toMsgbus)
	for {
		nc.Publish("arni", []byte(fmt.Sprintf("The time is now: %s", time.Now().Format(time.RFC3339))))
		time.Sleep(time.Second*3)
	}
}

func toMsgbus(msg *nats.Msg) {
	log.Println("from msgbus:", string(msg.Data))
}

func fromMsgbus(msg *nats.Msg) {
	log.Println("from msgbus:", string(msg.Data))
}
