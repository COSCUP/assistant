package main

import (
	"fmt"
	"github.com/PichuChen/daemon"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"
)

const (

	// name of the service
	name        = "coscup_assistant"
	description = "COSCUP Assistant"

	// port which daemon should be listen
	port = ":8171"
)

//	dependencies that are NOT required by the service, but might be used
var dependencies = []string{}

var stdlog, errlog *log.Logger

// Service has embedded daemon
type Service struct {
	daemon.Daemon
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {

	usage := "Usage: " + name + " install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	// Do something, call your goroutines, etc

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Set up listener for defined host and port
	// listener, err := net.Listen("tcp", port)

	go func() {
		http.HandleFunc("/", handler)
		err := http.ListenAndServe(port, nil)
		if err != nil {
			errlog.Println(err)
			// return "Possibly was a problem with the port binding", err
		}

	}()

	// set up channel on which to send accepted connections
	// listen := make(chan net.Conn, 100)
	// go acceptConnection(listener, listen)

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		// case conn := <-listen:
		// 	go handleClient(conn)
		case killSignal := <-interrupt:
			stdlog.Println("Got signal:", killSignal)
			// listener.Close()
			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}

	// never happen, but need to complete code
	return usage, nil
}

// Accept a client connection and collect it in a channel
func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

func handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		client.Write(buf[:numbytes])
	}
}

func init() {
	stdlog = log.New()
	errlog = log.New()

}

func main() {
	srv, err := daemon.New(name, description, dependencies...)
	if err != nil {
		errlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
