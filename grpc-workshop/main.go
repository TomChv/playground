package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"grpc-workshop/client"
	"grpc-workshop/server"
)

const errorMsg = "wrong arguments, you must tip '-c' to start the client or '-s' to start the server."

func exiError(msg string, code int) {
	log.Println(msg)
	os.Exit(code)
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		exiError(errorMsg, 1)
	}
	arg := args[0]

	if arg == "-s" {
		log.Println("Start the server")

		l, err := net.Listen("tcp", "localhost:9000")
		if err != nil {
			exiError(err.Error(), 1)
		}

		// Setup ^C channel
		quitChannel := make(chan os.Signal, 1)
		signal.Notify(quitChannel, os.Interrupt)

		serv := server.New(l)

		// Listen for ^C
		go func() {
			defer serv.Stop()
			_ = <-quitChannel
		}()

		err = serv.Start()
		if err != nil {
			exiError(err.Error(), 1)
		}

		return
	}

	if arg == "-c" {
		log.Println("Start the client")

		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			exiError(err.Error(), 1)
		}
		c := client.New(conn)
		res, err := c.Send(context.Background(), "Hello world")
		if err != nil {
			exiError(err.Error(), 1)
		}
		log.Println(fmt.Sprintf("Received response %v at %v", res.Body, res.Date))

		err = c.Disconnect()
		if err != nil {
			exiError(err.Error(), 1)
		}
		return
	}

	exiError(errorMsg, 1)
}
