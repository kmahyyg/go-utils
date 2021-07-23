package main

import (
	"flag"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

// TCP Port Scanner
var (
	flagThreads     int
	flagIPs         string
	flagResFile     string
	flagTimeout     int
	flagCustomPorts string

	baseScanner = &TCPScanner{}
)

func init() {
	flag.IntVar(&flagThreads, "t", runtime.NumCPU(), "Threads num")
	flag.StringVar(&flagResFile, "o", "result.txt", "Output File")
	flag.IntVar(&flagTimeout, "w", 2, "Timeout (seconds) ")
	flag.StringVar(&flagIPs, "i", "", "Scan IP Addresses, format: 192.168.1.1-192.168.1.222")
	flag.StringVar(&flagCustomPorts, "p", "", "Custom Ports, format: 21,22,23")
	flag.Parse()

	if len(flagIPs) < 7 {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func parsePortRange(s string) []int {
	portRange := make([]int, 0)
	if len(flagCustomPorts) == 0 {
		// merge default ports
		portRange = append(portRange, portWeb...)
		portRange = append(portRange, portRemote...)
		portRange = append(portRange, portFileTrans...)
		portRange = append(portRange, portDatabase...)
		portRange = append(portRange, portVirtualization...)
		portRange = append(portRange, portSpecial...)
		return portRange
	} else {
		tmpPortRange := String2IntSlice(flagCustomPorts)
		if len(tmpPortRange) == 0 {
			panic("Invalid Ports!")
		} else {
			return tmpPortRange
		}
	}
}

func parseIPRange(s string) []string {
	ipRange := strings.Split(s, "-")
	if len(ipRange) == 1 {
		testIP := net.ParseIP(ipRange[0])
		if testIP == nil {
			panic("Start IP Invalid!")
		}
		return ipRange
	} else if len(ipRange) > 2 {
		panic("Invalid IP Range.")
	} else {
		ipNetwork, err := ParseIPC(s)
		if err != nil {
			panic(err)
		}
		return ipNetwork
	}
}

func main() {
	// set max threads
	runtime.GOMAXPROCS(flagThreads*4 + 1)
	// check ulimit if too low which might lead to failure
	checkULimit()
	// save output bidirectionally
	fn := logOutput(flagResFile)
	defer fn()
	// parse input
	ipRange := parseIPRange(flagIPs)
	portRange := parsePortRange(flagCustomPorts)
	baseScanner.scanPorts = portRange
	baseScanner.scanNets = ipRange
	baseScanner.scanTimeout = time.Duration(flagTimeout) * time.Second
	baseScanner.scanThreads = flagThreads
	// start scan
	runtime.GC()
	baseScanner.Scan()
}
