package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func checkdomain(domain string) {
	fmt.Println("\nChecking Domain:", domain)

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// MX Records
	mxrecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Printf("MX lookup failed: %v\n", err)
	} else if len(mxrecords) > 0 {
		hasMX = true
	}

	// SPF Records
	txtrecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Printf("SPF lookup failed: %v\n", err)
	} else {
		for _, r := range txtrecords {
			if strings.HasPrefix(r, "v=spf1") {
				hasSPF = true
				spfRecord = r
				break
			}
		}
	}

	// DMARC Records
	dmarcrecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Printf("DMARC lookup failed: %v\n", err)
	} else {
		for _, r := range dmarcrecords {
			if strings.HasPrefix(r, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = r
				break
			}
		}
	}

	
	if hasMX {
		fmt.Println(" MX record found")
	} else {
		fmt.Println(" No MX record found")
	}

	if hasSPF {
		fmt.Printf(" SPF record found: %v\n", spfRecord)
	} else {
		fmt.Println(" No SPF record found")
	}

	if hasDMARC {
		fmt.Printf(" DMARC record found: %v\n", dmarcRecord)
	} else {
		fmt.Println(" No DMARC record found")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your Domains to check:")

	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain != "" {
			checkdomain(domain)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Input error: %v\n", err)
	}
}
