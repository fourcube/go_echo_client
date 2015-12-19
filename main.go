// A simple echo client written in go
package main // import "github.com/fourcube/go_echo_client"

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = start

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "echo_host",
			Value:  "127.0.0.1",
			EnvVar: "ECHO_HOST",
		},
		cli.StringFlag{
			Name:   "echo_port",
			Value:  "4040",
			EnvVar: "ECHO_PORT",
		},
	}

	app.Run(os.Args)
}

func start(c *cli.Context) {
	echoService := fmt.Sprintf("%s:%s", c.String("echo_host"), c.String("echo_port"))

	conn, err := net.Dial("tcp", echoService)
	if err != nil {
		log.Fatalf("Couldn't connect to echo service at %v: %v", echoService, err)
	}

	reader := bufio.NewReader(conn)

	for {
		msg := fmt.Sprintf("hi\n")

		fmt.Printf("%s -> %v\n", strings.TrimSpace(string(msg)), conn.RemoteAddr())
		bytesWritten, err := conn.Write([]byte(msg))
		if err != nil {
			log.Printf("Failed write on tcp socket: %v", err)
			return
		}

		if bytesWritten != len(msg) {
			log.Printf("Warning: Incomplete write on tcp socket")
			return
		}

		data, err := reader.ReadBytes('\n')
		if err != nil {
			log.Printf("Couldn't read from tcp socket: %v", err)
			return
		}

		fmt.Printf("%s <- %v\n", strings.TrimSpace(string(data)), conn.RemoteAddr())

		time.Sleep(5 * time.Second)
	}
}
