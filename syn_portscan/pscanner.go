package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func connectTCP(targets chan string, timeout time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	for target := range targets {
		if len(target) < 2 {
			return
		}
		conn, err := net.DialTimeout("tcp", target, timeout)
		if err != nil || conn == nil {
			if strings.Contains(err.Error(), "too many open files") {
				time.Sleep(time.Duration(flagTimeout) * time.Second)
			} else {
				fmt.Fprintf(os.Stderr, "%s CLOSED! \n", target)
			}
		} else {
			defer conn.Close()
			log.Printf("%s OPEN! \n", target)
		}
	}
}

func (pscanner *TCPScanner) Scan() {
	var wg = sync.WaitGroup{}
	var targets = make(chan string, runtime.NumCPU()*4)
	for t := 0; t < pscanner.scanThreads; t++ {
		wg.Add(1)
		go connectTCP(targets, pscanner.scanTimeout, &wg)
	}
	for _, vport := range pscanner.scanPorts {
		for _, vip := range pscanner.scanNets {
			target := fmt.Sprintf("%s:%d", vip, vport)
			targets <- target
		}
	}
	close(targets)
	wg.Wait()
}
