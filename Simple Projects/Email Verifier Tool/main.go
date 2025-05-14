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

	// MX
	mxrecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Printf("MX lookup failed: %v\n", err)
	} else if len(mxrecords) > 0 {
		hasMX = true
	}

	// SPF
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

	// DMARC
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

func readFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text()) // line splitting
		if domain != "" {
			domains = append(domains, domain)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	return domains
}

func readFromCLI() []string {
	fmt.Print("Enter space-separated domains: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		domains := strings.Fields(line) // space splitting
		return domains
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	return nil
}

func main() {
	var domains []string

	if len(os.Args) == 2 {
		domains = readFromFile(os.Args[1])
	} else {
		domains = readFromCLI()
	}

	fmt.Println("\n--- Results ---")
	for _, domain := range domains {
		checkdomain(domain)
	}
}
