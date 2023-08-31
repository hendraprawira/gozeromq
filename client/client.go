// Copyright 2018 The go-zeromq Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// PubSub envelope publisher
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-zeromq/zmq4"
)

func main() {
	log.SetPrefix("psenvpub: ")

	// prepare the publisher
	pub := zmq4.NewPub(context.Background())
	defer pub.Close()

	err := pub.Listen("tcp://*:8001")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	msgA := zmq4.NewMsgFrom(
		[]byte("A"),
		[]byte("We don't want to see this"),
	)
	msgB := zmq4.NewMsgFrom(
		[]byte("B"),
		[]byte("We would like to see this"),
	)
	for {
		//  Write two messages, each with an envelope and content
		err = pub.Send(msgA)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(msgA, " message has been sent")
		}
		err = pub.Send(msgB)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(msgB, " message has been sent")
		}
		time.Sleep(time.Second)
	}
}

// ignore
// // +build ignore

// package p // Copyright 2018 The go-zeromq Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// // license that can be found in the LICENSE file.

// //go:build ignore
// // +build ignore

// // PubSub envelope publisher
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	zmq4 "github.com/go-zeromq/zmq4"
// )

// var byteSlice []zmq4.Msg

// func main() {
// 	log.SetPrefix("psenvpub: ")

// 	// prepare the publisher
// 	pub := zmq4.NewPush(context.Background(), zmq4.WithDialerMaxRetries(-1))
// 	defer pub.Close()

// 	err := pub.Listen("tcp://*:5563")
// 	if err != nil {
// 		log.Fatalf("could not listen: %v", err)
// 	}

// 	// msgA := zmq4.NewMsgFrom(
// 	// 	[]byte("A"),
// 	// 	[]byte("We don't want to see this"),
// 	// )
// 	// msgB := zmq4.NewMsgFrom(
// 	// 	[]byte("B"),
// 	// 	[]byte("We would like to see this"),
// 	// )
// 	i := 0

// 	for {
// 		//  Write two messages, each with an envelope and content
// 		msgA := zmq4.NewMsgFrom(
// 			[]byte("B"),
// 			[]byte(fmt.Sprint(i)),
// 		)
// 		// _ = pub.Send(msgA)
// 		sendMessages(msgA, pub)
// 		// if err != nil {
// 		// log.Fatal(err)
// 		// }
// 		i++
// 		// err = pub.Send(msgB)
// 		// if err != nil {
// 		// 	log.Fatal(err)
// 		// }
// 		time.Sleep(time.Second)
// 	}
// }

// func sendMessages(msg zmq4.Msg, pub zmq4.Socket) {
// 	byteSlice = append(byteSlice, msg)
// 	if len(byteSlice) == 0 {
// 		for _, value := range byteSlice {
// 			err := pub.Send(value)
// 			if err != nil {
// 				// log.Fatal(err)
// 				byteSlice = append(byteSlice, msg)
// 			} else {
// 				byteSlice = byteSlice[:len(byteSlice)-1]

// 			}
// 		}

// 	}
// }

// Copyright 2018 The go-zeromq Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// PubSub envelope publisher
// package main

// import (
// 	"container/list"
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	zmq4 "github.com/go-zeromq/zmq4"
// )

// func main() {
// 	pub := zmq4.NewPush(context.Background(), zmq4.WithAutomaticReconnect(true))
// 	defer pub.Close()
// 	err1 := pub.Listen("tcp://*:5563")
// 	if err1 != nil {
// 		log.Fatalf("could not listen: %v", err1)
// 	}
// 	messageBuffer := list.New()
// 	// var messageBufferError [1]int
// 	i := 0
// 	// var errs error
// 	for {
// 		// log.Print(err1)
// 		message := fmt.Sprintf("%04d", i)
// 		messageBuffer.PushBack(message)
// 		for messageBuffer.Len() > 0 {
// 			element := messageBuffer.Front()
// 			messageBytes := []byte(element.Value.(string))
// 			zmqMessage := zmq4.NewMsg(messageBytes)
// 			// if errs != nil {
// 			errsd := pub.Send(zmqMessage)
// 			// log.Print(pub.Close().Error(), errs)
// 			fmt.Print(errsd)
// 			fmt.Println("Pushed:", element.Value)
// 			messageBuffer.Remove(element)
// 			// } else {
// 			// 	messageBufferError[0] = 1
// 			// 	// fmt.Print(len(messageBufferError))
// 			// }
// 		}
// 		// log.Print(errs, len(messageBufferError))
// 		// fmt.Println(message)
// 		i++
// 		time.Sleep(time.Second)
// 	}
// }
