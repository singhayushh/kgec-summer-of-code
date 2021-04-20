package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/rcrowley/goagain"
)

var s = true

func HandleInterrupt() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(char)

	switch char {
	case 'x':
		s =
			main()
	case 's':
		s = false
	case 'a':
		fmt.Println("a Key Pressed")
	case 'A':
		fmt.Println("terminating....")
		os.Exit(3)
	}
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix(fmt.Sprintf("pid:%d ", syscall.Getpid()))
}

func main() {
	go func() {
		for {
			HandleInterrupt()
		}
	}()
	// Inherit a net.Listener from our parent process or listen anew.
	l, err := goagain.Listener()
	if s {

		// Listen on a TCP or a UNIX domain socket (TCP here).
		l, err = net.Listen("tcp", "127.0.0.1:48879")
		if nil != err {
			log.Fatalln(err)
		}
		log.Println("listening on", l.Addr())

		// Accept connections in a new goroutine.
		go serve(l)

	} else {

		// Resume accepting connections in a new goroutine.
		log.Println("resuming listening on", l.Addr())
		go serve(l)

		// Kill the parent, now that the child has started successfully.
		if err := goagain.Kill(); nil != err {
			log.Fatalln(err)
		}

	}

	// Block the main goroutine awaiting signals.
	if _, err := goagain.Wait(l); nil != err {
		log.Fatalln(err)
	}

	// Do whatever's necessary to ensure a graceful exit like waiting for
	// goroutines to terminate or a channel to become closed.
	//
	// In this case, we'll simply stop listening and wait one second.
	if err := l.Close(); nil != err {
		log.Fatalln(err)
	}
	time.Sleep(1e9)

}

// A very rude server that says hello and then closes your connection.
func serve(l net.Listener) {
	for {
		c, err := l.Accept()
		if nil != err {
			if goagain.IsErrClosing(err) {
				break
			}
			log.Fatalln(err)
		}
		c.Write([]byte("Hello, world!\n"))
		c.Close()
	}
}
