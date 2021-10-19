package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/koron-go/sigctx"
)

var enableSpawn = true

func main() {
	flag.BoolVar(&enableSpawn, "spawn", true, "")
	flag.Parse()
	ctx, cancel := sigctx.WithCancelSignal(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	log.Printf("The program is waiting SIGINT or SIGTERM. Press Ctrl-C")
	<-ctx.Done()
	log.Printf("recived signal (Go)")
	if enableSpawn {
		cmd := exec.Command("echo", "received signal (echo)")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Printf("something wrong: %v", err)
			return
		}
	}
	log.Printf("completed")
}
