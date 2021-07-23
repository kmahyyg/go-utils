package main

import (
	"time"
)

type TCPScanner struct {
	scanNets    []string
	scanPorts   []int
	scanTimeout time.Duration
	scanThreads int
}

var (
	portWeb            = []int{443, 80, 8080, 7001, 7002, 9060, 9080, 9443, 8443, 3000, 9000, 9090}
	portRemote         = []int{22, 5938, 5985, 3389, 1080, 5800, 5900}
	portFileTrans      = []int{21, 23, 139, 445, 135, 2121, 2049, 3690}
	portDatabase       = []int{3306, 27017, 1433, 1521, 61616, 6379, 9200, 15672, 5432}
	portVirtualization = []int{902, 903, 2375, 5000}
	portSpecial        = []int{8090, 8009, 4430, 7012, 8088, 18080, 28080}
)
