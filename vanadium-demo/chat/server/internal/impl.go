package internal

import (
	"bufio"
	"log"
	"os"
	"strings"

	"se.uom.gr/chat"
	"v.io/v23/context"
	"v.io/v23/rpc"
)

type impl struct {
}

func NewImpl() chat.ChatServerMethods {
	return &impl{}
}

// SendMessage sends a new message
func (f *impl) SendMessage(_ *context.T, _ rpc.ServerCall, msg string) (string, error) {
	log.Printf("%v", msg)
	reply := getInput()
	log.Printf("%v", reply)
	return reply, nil
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSuffix(msg, "\n")
	return msg
}
