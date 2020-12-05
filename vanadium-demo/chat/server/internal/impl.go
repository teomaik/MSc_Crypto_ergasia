// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"log"
	"math/rand"
	"sync"

	"se.uom.gr/chat"
	"v.io/v23/context"
	"v.io/v23/rpc"
)

// Fortune implements fortune.FortuneServerMethods.
type impl struct {
	fortunes []string
	random   *rand.Rand
	mutex    sync.RWMutex
}

// newImpl returns a Fortune implementation that can be used to create a service.
func NewImpl() chat.ChatServerMethods {
	return &impl{
		fortunes: []string{
			"You will reach the heights of success.",
			"Conquer your fears or they will conquer you.",
			"Today is your lucky day!",
		},
		random: rand.New(rand.NewSource(99)),
	}
}

// SendMessage sends a new message to other client.
func (f *impl) SendMessage(_ *context.T, _ rpc.ServerCall, msg string) error {
	log.Println(msg)
	return nil
}
