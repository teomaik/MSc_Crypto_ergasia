// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command fortune is a client to the Fortune interface.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	v23 "v.io/v23"
	"v.io/v23/context"
	_ "v.io/x/ref/runtime/factories/generic"

	"se.uom.gr/chat"
)

var (
	name = flag.String("name", "/169.254.218.122:1446", "Name of the server to connect to")
	msg  string
)

func main() {

	ctx, shutdown := v23.Init()
	defer shutdown()

	client := chat.ChatClient(*name)

	// Create a child context that will timeout in 60s.
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	msg, _ := reader.ReadString('\n')

	switch {
	case msg != "":
		err := client.SendMessage(ctx, msg)
		if err != nil {
			log.Panic("Error sending message: ", err)
		}
	}
}
