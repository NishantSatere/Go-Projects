package main

import (
	"bufio"
	"fmt"
	// "go/scanner"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(Domain string) {
	hasMx := false
	hasDmarc := false
	hashSPF := false
	spfRecord := ""
	dmarcRecord := ""

	mxs, err := net.LookupMX(Domain)

	if err != nil {
		fmt.Println(err)
	}

	if len(mxs) > 0 {
		hasMx = true
	}

	txt, err := net.LookupTXT(Domain)

	if err != nil {
		fmt.Println(err)
	}

	for _, txtRecord := range txt {
		if strings.HasPrefix(txtRecord, "v=spf1") {
			hashSPF = true
			spfRecord = txtRecord
			break
		}
	}

	dmarc, err := net.LookupTXT("_dmarc." + Domain)
	if err != nil {
		fmt.Println(err)
	}

	for _, dmarcRecord = range dmarc {
		if strings.HasPrefix(dmarcRecord, "v=DMARC1") {
			hasDmarc = true
			break
		}
	}

	fmt.Printf("%s, %t, %t, %t, %s, %t, %s\n", Domain, hasMx, hashSPF, spfRecord, hasDmarc, dmarcRecord)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMx, hashSPF, spfRecord, hasDmarc, dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
