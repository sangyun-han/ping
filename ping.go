package main

import (
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"fmt"
	"time"
)

func main() {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("ping from %s, time=%v\n", addr.String(), rtt)
	}

	for {
		err = p.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

