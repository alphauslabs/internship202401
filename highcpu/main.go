package main

import (
	"flag"
	"log"
	"runtime"
	"time"
)

var (
	multi = flag.Bool("multi", false, "Set to true to use all CPU/cores")
)

func main() {
	flag.Parse()
	numCpu := 1
	if *multi {
		numCpu = runtime.NumCPU()
	}

	done := make(chan int)
	for i := 0; i < numCpu; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		if !*multi {
			log.Println("loading a single CPU...")
		} else {
			log.Println("loading all CPUs...")
		}

		time.Sleep(time.Second * 1)
	}

	close(done)
}
