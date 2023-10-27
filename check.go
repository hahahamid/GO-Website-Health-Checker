package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[THIS DOMAIN IS DOWN] \n %v is unreachable,\n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("\n[LESSGGOOO THE DOMAIN IS UP]\n %v is reachable,\n\n From: %v\n To: %v", destination,
			conn.LocalAddr(),
			conn.RemoteAddr())
	}
	return status
}
