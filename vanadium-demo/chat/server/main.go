package main

import (
	"flag"
	"log"

	"se.uom.gr/chat"
	"se.uom.gr/chat/server/internal"
	v23 "v.io/v23"
	"v.io/v23/security"
	"v.io/x/ref/lib/signals"

	_ "v.io/x/ref/runtime/factories/roaming"
)

var (
	name = flag.String("name", "", "Name for fortuned in default mount table")
)

func main() {
	ctx, shutdown := v23.Init()
	defer shutdown()

	authorizer := security.DefaultAuthorizer()
	impl := internal.NewImpl()
	service := chat.ChatServer(impl)
	ctx, server, err := v23.WithNewServer(ctx, *name, service, authorizer)
	if err != nil {
		log.Panic("Failure creating server: ", err)
	}
	log.Printf("Listening at: %v\n\n\n", server.Status().Endpoints[0])

	<-signals.ShutdownOnSignals(ctx)
}
