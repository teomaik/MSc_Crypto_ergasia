package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	v23 "v.io/v23"
	"v.io/v23/context"
	_ "v.io/x/ref/runtime/factories/generic"

	"se.uom.gr/chat"
)

var (
	serverName = flag.String("serverName", "/127.0.0.1:2507", "Name of the server to connect to")
)

func main() {

	ctx, shutdown := v23.Init()
	defer shutdown()

	client := chat.ChatClient(*serverName)

	/* Create a child context that will timeout in 60s. */
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	for true {
		var msg string = getInput()
		value, _ := client.SendMessage(ctx, msg)
		log.Printf("%v", value)
	}

}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSuffix(msg, "\n")
	return msg
}
