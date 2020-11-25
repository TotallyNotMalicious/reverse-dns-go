package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := flag.Bool("r", false, "Reverse DNS")

	flag.Parse()
	host := flag.Args()

		if *arguments {

			addresses, err := net.LookupAddr(host[0])
			if err != nil {
				fmt.Printf("\nError: %v\n", err)
				os.Exit(1)
			}
			for _, address := range addresses {
				fmt.Printf("\n")
				fmt.Printf(host[0]+" Has Resolved To %s\n", address)
			}
		} else {
			ips, err := net.LookupIP(host[0])
			if err != nil {
				fmt.Printf("\n")
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, ip := range ips {
				fmt.Printf("\n")
				fmt.Printf(host[0]+" Has Resolved To %s\n", ip)
			}
		}
	}